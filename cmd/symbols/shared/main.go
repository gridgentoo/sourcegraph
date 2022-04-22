package shared

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/inconshreveable/log15"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/sourcegraph/sourcegraph/cmd/symbols/fetcher"
	"github.com/sourcegraph/sourcegraph/cmd/symbols/gitserver"
	"github.com/sourcegraph/sourcegraph/cmd/symbols/internal/api"
	"github.com/sourcegraph/sourcegraph/cmd/symbols/types"
	"github.com/sourcegraph/sourcegraph/internal/actor"
	"github.com/sourcegraph/sourcegraph/internal/conf"
	"github.com/sourcegraph/sourcegraph/internal/debugserver"
	"github.com/sourcegraph/sourcegraph/internal/env"
	"github.com/sourcegraph/sourcegraph/internal/goroutine"
	"github.com/sourcegraph/sourcegraph/internal/honey"
	"github.com/sourcegraph/sourcegraph/internal/httpserver"
	"github.com/sourcegraph/sourcegraph/internal/logging"
	"github.com/sourcegraph/sourcegraph/internal/observation"
	"github.com/sourcegraph/sourcegraph/internal/profiler"
	"github.com/sourcegraph/sourcegraph/internal/sentry"
	"github.com/sourcegraph/sourcegraph/internal/trace"
	"github.com/sourcegraph/sourcegraph/internal/trace/ot"
	"github.com/sourcegraph/sourcegraph/internal/tracer"
)

const addr = ":3184"

type SetupFunc func(observationContext *observation.Context, gitserverClient gitserver.GitserverClient, repositoryFetcher fetcher.RepositoryFetcher) (types.SearchFunc, func(http.ResponseWriter, *http.Request), []goroutine.BackgroundRoutine, string, error)

func Main(setup SetupFunc) {
	routines := []goroutine.BackgroundRoutine{}

	// Initialize tracing/metrics
	observationContext := &observation.Context{
		Logger:     log15.Root(),
		Tracer:     &trace.Tracer{Tracer: opentracing.GlobalTracer()},
		Registerer: prometheus.DefaultRegisterer,
		HoneyDataset: &honey.Dataset{
			Name:       "codeintel-symbols",
			SampleRate: 5,
		},
	}

	// Conf package must be initialized prior to Rockskip init
	conf.Init()

	// Run setup
	gitserverClient := gitserver.NewClient(observationContext)
	repositoryFetcher := fetcher.NewRepositoryFetcher(gitserverClient, types.LoadRepositoryFetcherConfig(env.BaseConfig{}).MaxTotalPathsLength, observationContext)
	searchFunc, handleStatus, newRoutines, ctagsBinary, err := setup(observationContext, gitserverClient, repositoryFetcher)
	if err != nil {
		log.Fatalf("Failed to setup: %v", err)
	}
	routines = append(routines, newRoutines...)

	// Initialization
	env.HandleHelpFlag()
	logging.Init()
	tracer.Init(conf.DefaultClient())
	sentry.Init(conf.DefaultClient())
	trace.Init()
	profiler.Init()

	// Start debug server
	ready := make(chan struct{})
	go debugserver.NewServerRoutine(ready).Start()

	// Create HTTP server
	server := httpserver.NewFromAddr(addr, &http.Server{
		ReadTimeout:  75 * time.Second,
		WriteTimeout: 10 * time.Minute,
		Handler:      actor.HTTPMiddleware(ot.HTTPMiddleware(trace.HTTPMiddleware(api.NewHandler(searchFunc, handleStatus, ctagsBinary), conf.DefaultClient()))),
	})
	routines = append(routines, server)

	// Mark health server as ready and go!
	close(ready)
	goroutine.MonitorBackgroundRoutines(context.Background(), routines...)
}
