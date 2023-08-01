package accessrequests

import (
	"context"
	"sync"

	database "github.com/sourcegraph/sourcegraph/internal/database"
	basestore "github.com/sourcegraph/sourcegraph/internal/database/basestore"
	types "github.com/sourcegraph/sourcegraph/internal/types"
)

// MockAccessRequestStore is a mock implementation of the AccessRequestStore
// interface (from the package
// github.com/sourcegraph/sourcegraph/internal/database) used for unit
// testing.
type MockAccessRequestStore struct {
	// CountFunc is an instance of a mock function object controlling the
	// behavior of the method Count.
	CountFunc *AccessRequestStoreCountFunc
	// CreateFunc is an instance of a mock function object controlling the
	// behavior of the method Create.
	CreateFunc *AccessRequestStoreCreateFunc
	// DoneFunc is an instance of a mock function object controlling the
	// behavior of the method Done.
	DoneFunc *AccessRequestStoreDoneFunc
	// GetByEmailFunc is an instance of a mock function object controlling
	// the behavior of the method GetByEmail.
	GetByEmailFunc *AccessRequestStoreGetByEmailFunc
	// GetByIDFunc is an instance of a mock function object controlling the
	// behavior of the method GetByID.
	GetByIDFunc *AccessRequestStoreGetByIDFunc
	// HandleFunc is an instance of a mock function object controlling the
	// behavior of the method Handle.
	HandleFunc *AccessRequestStoreHandleFunc
	// ListFunc is an instance of a mock function object controlling the
	// behavior of the method List.
	ListFunc *AccessRequestStoreListFunc
	// UpdateFunc is an instance of a mock function object controlling the
	// behavior of the method Update.
	UpdateFunc *AccessRequestStoreUpdateFunc
	// WithTransactFunc is an instance of a mock function object controlling
	// the behavior of the method WithTransact.
	WithTransactFunc *AccessRequestStoreWithTransactFunc
}

// NewMockAccessRequestStore creates a new mock of the AccessRequestStore
// interface. All methods return zero values for all results, unless
// overwritten.
func NewMockAccessRequestStore() *MockAccessRequestStore {
	return &MockAccessRequestStore{
		CountFunc: &AccessRequestStoreCountFunc{
			defaultHook: func(context.Context, *AccessRequestsFilterArgs) (r0 int, r1 error) {
				return
			},
		},
		CreateFunc: &AccessRequestStoreCreateFunc{
			defaultHook: func(context.Context, *types.AccessRequest) (r0 *types.AccessRequest, r1 error) {
				return
			},
		},
		DoneFunc: &AccessRequestStoreDoneFunc{
			defaultHook: func(error) (r0 error) {
				return
			},
		},
		GetByEmailFunc: &AccessRequestStoreGetByEmailFunc{
			defaultHook: func(context.Context, string) (r0 *types.AccessRequest, r1 error) {
				return
			},
		},
		GetByIDFunc: &AccessRequestStoreGetByIDFunc{
			defaultHook: func(context.Context, int32) (r0 *types.AccessRequest, r1 error) {
				return
			},
		},
		HandleFunc: &AccessRequestStoreHandleFunc{
			defaultHook: func() (r0 basestore.TransactableHandle) {
				return
			},
		},
		ListFunc: &AccessRequestStoreListFunc{
			defaultHook: func(context.Context, *AccessRequestsFilterArgs, *database.PaginationArgs) (r0 []*types.AccessRequest, r1 error) {
				return
			},
		},
		UpdateFunc: &AccessRequestStoreUpdateFunc{
			defaultHook: func(context.Context, *types.AccessRequest) (r0 *types.AccessRequest, r1 error) {
				return
			},
		},
		WithTransactFunc: &AccessRequestStoreWithTransactFunc{
			defaultHook: func(context.Context, func(AccessRequestStore) error) (r0 error) {
				return
			},
		},
	}
}

// NewStrictMockAccessRequestStore creates a new mock of the
// AccessRequestStore interface. All methods panic on invocation, unless
// overwritten.
func NewStrictMockAccessRequestStore() *MockAccessRequestStore {
	return &MockAccessRequestStore{
		CountFunc: &AccessRequestStoreCountFunc{
			defaultHook: func(context.Context, *AccessRequestsFilterArgs) (int, error) {
				panic("unexpected invocation of MockAccessRequestStore.Count")
			},
		},
		CreateFunc: &AccessRequestStoreCreateFunc{
			defaultHook: func(context.Context, *types.AccessRequest) (*types.AccessRequest, error) {
				panic("unexpected invocation of MockAccessRequestStore.Create")
			},
		},
		DoneFunc: &AccessRequestStoreDoneFunc{
			defaultHook: func(error) error {
				panic("unexpected invocation of MockAccessRequestStore.Done")
			},
		},
		GetByEmailFunc: &AccessRequestStoreGetByEmailFunc{
			defaultHook: func(context.Context, string) (*types.AccessRequest, error) {
				panic("unexpected invocation of MockAccessRequestStore.GetByEmail")
			},
		},
		GetByIDFunc: &AccessRequestStoreGetByIDFunc{
			defaultHook: func(context.Context, int32) (*types.AccessRequest, error) {
				panic("unexpected invocation of MockAccessRequestStore.GetByID")
			},
		},
		HandleFunc: &AccessRequestStoreHandleFunc{
			defaultHook: func() basestore.TransactableHandle {
				panic("unexpected invocation of MockAccessRequestStore.Handle")
			},
		},
		ListFunc: &AccessRequestStoreListFunc{
			defaultHook: func(context.Context, *AccessRequestsFilterArgs, *database.PaginationArgs) ([]*types.AccessRequest, error) {
				panic("unexpected invocation of MockAccessRequestStore.List")
			},
		},
		UpdateFunc: &AccessRequestStoreUpdateFunc{
			defaultHook: func(context.Context, *types.AccessRequest) (*types.AccessRequest, error) {
				panic("unexpected invocation of MockAccessRequestStore.Update")
			},
		},
		WithTransactFunc: &AccessRequestStoreWithTransactFunc{
			defaultHook: func(context.Context, func(AccessRequestStore) error) error {
				panic("unexpected invocation of MockAccessRequestStore.WithTransact")
			},
		},
	}
}

// NewMockAccessRequestStoreFrom creates a new mock of the
// MockAccessRequestStore interface. All methods delegate to the given
// implementation, unless overwritten.
func NewMockAccessRequestStoreFrom(i AccessRequestStore) *MockAccessRequestStore {
	return &MockAccessRequestStore{
		CountFunc: &AccessRequestStoreCountFunc{
			defaultHook: i.Count,
		},
		CreateFunc: &AccessRequestStoreCreateFunc{
			defaultHook: i.Create,
		},
		DoneFunc: &AccessRequestStoreDoneFunc{
			defaultHook: i.Done,
		},
		GetByEmailFunc: &AccessRequestStoreGetByEmailFunc{
			defaultHook: i.GetByEmail,
		},
		GetByIDFunc: &AccessRequestStoreGetByIDFunc{
			defaultHook: i.GetByID,
		},
		HandleFunc: &AccessRequestStoreHandleFunc{
			defaultHook: i.Handle,
		},
		ListFunc: &AccessRequestStoreListFunc{
			defaultHook: i.List,
		},
		UpdateFunc: &AccessRequestStoreUpdateFunc{
			defaultHook: i.Update,
		},
		WithTransactFunc: &AccessRequestStoreWithTransactFunc{
			defaultHook: i.WithTransact,
		},
	}
}

// AccessRequestStoreCountFunc describes the behavior when the Count method
// of the parent MockAccessRequestStore instance is invoked.
type AccessRequestStoreCountFunc struct {
	defaultHook func(context.Context, *AccessRequestsFilterArgs) (int, error)
	hooks       []func(context.Context, *AccessRequestsFilterArgs) (int, error)
	history     []AccessRequestStoreCountFuncCall
	mutex       sync.Mutex
}

// Count delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockAccessRequestStore) Count(v0 context.Context, v1 *AccessRequestsFilterArgs) (int, error) {
	r0, r1 := m.CountFunc.nextHook()(v0, v1)
	m.CountFunc.appendCall(AccessRequestStoreCountFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the Count method of the
// parent MockAccessRequestStore instance is invoked and the hook queue is
// empty.
func (f *AccessRequestStoreCountFunc) SetDefaultHook(hook func(context.Context, *AccessRequestsFilterArgs) (int, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Count method of the parent MockAccessRequestStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *AccessRequestStoreCountFunc) PushHook(hook func(context.Context, *AccessRequestsFilterArgs) (int, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *AccessRequestStoreCountFunc) SetDefaultReturn(r0 int, r1 error) {
	f.SetDefaultHook(func(context.Context, *AccessRequestsFilterArgs) (int, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *AccessRequestStoreCountFunc) PushReturn(r0 int, r1 error) {
	f.PushHook(func(context.Context, *AccessRequestsFilterArgs) (int, error) {
		return r0, r1
	})
}

func (f *AccessRequestStoreCountFunc) nextHook() func(context.Context, *AccessRequestsFilterArgs) (int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *AccessRequestStoreCountFunc) appendCall(r0 AccessRequestStoreCountFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of AccessRequestStoreCountFuncCall objects
// describing the invocations of this function.
func (f *AccessRequestStoreCountFunc) History() []AccessRequestStoreCountFuncCall {
	f.mutex.Lock()
	history := make([]AccessRequestStoreCountFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// AccessRequestStoreCountFuncCall is an object that describes an invocation
// of method Count on an instance of MockAccessRequestStore.
type AccessRequestStoreCountFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 *AccessRequestsFilterArgs
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 int
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c AccessRequestStoreCountFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c AccessRequestStoreCountFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// AccessRequestStoreCreateFunc describes the behavior when the Create
// method of the parent MockAccessRequestStore instance is invoked.
type AccessRequestStoreCreateFunc struct {
	defaultHook func(context.Context, *types.AccessRequest) (*types.AccessRequest, error)
	hooks       []func(context.Context, *types.AccessRequest) (*types.AccessRequest, error)
	history     []AccessRequestStoreCreateFuncCall
	mutex       sync.Mutex
}

// Create delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockAccessRequestStore) Create(v0 context.Context, v1 *types.AccessRequest) (*types.AccessRequest, error) {
	r0, r1 := m.CreateFunc.nextHook()(v0, v1)
	m.CreateFunc.appendCall(AccessRequestStoreCreateFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the Create method of the
// parent MockAccessRequestStore instance is invoked and the hook queue is
// empty.
func (f *AccessRequestStoreCreateFunc) SetDefaultHook(hook func(context.Context, *types.AccessRequest) (*types.AccessRequest, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Create method of the parent MockAccessRequestStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *AccessRequestStoreCreateFunc) PushHook(hook func(context.Context, *types.AccessRequest) (*types.AccessRequest, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *AccessRequestStoreCreateFunc) SetDefaultReturn(r0 *types.AccessRequest, r1 error) {
	f.SetDefaultHook(func(context.Context, *types.AccessRequest) (*types.AccessRequest, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *AccessRequestStoreCreateFunc) PushReturn(r0 *types.AccessRequest, r1 error) {
	f.PushHook(func(context.Context, *types.AccessRequest) (*types.AccessRequest, error) {
		return r0, r1
	})
}

func (f *AccessRequestStoreCreateFunc) nextHook() func(context.Context, *types.AccessRequest) (*types.AccessRequest, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *AccessRequestStoreCreateFunc) appendCall(r0 AccessRequestStoreCreateFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of AccessRequestStoreCreateFuncCall objects
// describing the invocations of this function.
func (f *AccessRequestStoreCreateFunc) History() []AccessRequestStoreCreateFuncCall {
	f.mutex.Lock()
	history := make([]AccessRequestStoreCreateFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// AccessRequestStoreCreateFuncCall is an object that describes an
// invocation of method Create on an instance of MockAccessRequestStore.
type AccessRequestStoreCreateFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 *types.AccessRequest
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *types.AccessRequest
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c AccessRequestStoreCreateFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c AccessRequestStoreCreateFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// AccessRequestStoreDoneFunc describes the behavior when the Done method of
// the parent MockAccessRequestStore instance is invoked.
type AccessRequestStoreDoneFunc struct {
	defaultHook func(error) error
	hooks       []func(error) error
	history     []AccessRequestStoreDoneFuncCall
	mutex       sync.Mutex
}

// Done delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockAccessRequestStore) Done(v0 error) error {
	r0 := m.DoneFunc.nextHook()(v0)
	m.DoneFunc.appendCall(AccessRequestStoreDoneFuncCall{v0, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Done method of the
// parent MockAccessRequestStore instance is invoked and the hook queue is
// empty.
func (f *AccessRequestStoreDoneFunc) SetDefaultHook(hook func(error) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Done method of the parent MockAccessRequestStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *AccessRequestStoreDoneFunc) PushHook(hook func(error) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *AccessRequestStoreDoneFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(error) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *AccessRequestStoreDoneFunc) PushReturn(r0 error) {
	f.PushHook(func(error) error {
		return r0
	})
}

func (f *AccessRequestStoreDoneFunc) nextHook() func(error) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *AccessRequestStoreDoneFunc) appendCall(r0 AccessRequestStoreDoneFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of AccessRequestStoreDoneFuncCall objects
// describing the invocations of this function.
func (f *AccessRequestStoreDoneFunc) History() []AccessRequestStoreDoneFuncCall {
	f.mutex.Lock()
	history := make([]AccessRequestStoreDoneFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// AccessRequestStoreDoneFuncCall is an object that describes an invocation
// of method Done on an instance of MockAccessRequestStore.
type AccessRequestStoreDoneFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 error
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c AccessRequestStoreDoneFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c AccessRequestStoreDoneFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// AccessRequestStoreGetByEmailFunc describes the behavior when the
// GetByEmail method of the parent MockAccessRequestStore instance is
// invoked.
type AccessRequestStoreGetByEmailFunc struct {
	defaultHook func(context.Context, string) (*types.AccessRequest, error)
	hooks       []func(context.Context, string) (*types.AccessRequest, error)
	history     []AccessRequestStoreGetByEmailFuncCall
	mutex       sync.Mutex
}

// GetByEmail delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockAccessRequestStore) GetByEmail(v0 context.Context, v1 string) (*types.AccessRequest, error) {
	r0, r1 := m.GetByEmailFunc.nextHook()(v0, v1)
	m.GetByEmailFunc.appendCall(AccessRequestStoreGetByEmailFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetByEmail method of
// the parent MockAccessRequestStore instance is invoked and the hook queue
// is empty.
func (f *AccessRequestStoreGetByEmailFunc) SetDefaultHook(hook func(context.Context, string) (*types.AccessRequest, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetByEmail method of the parent MockAccessRequestStore instance invokes
// the hook at the front of the queue and discards it. After the queue is
// empty, the default hook function is invoked for any future action.
func (f *AccessRequestStoreGetByEmailFunc) PushHook(hook func(context.Context, string) (*types.AccessRequest, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *AccessRequestStoreGetByEmailFunc) SetDefaultReturn(r0 *types.AccessRequest, r1 error) {
	f.SetDefaultHook(func(context.Context, string) (*types.AccessRequest, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *AccessRequestStoreGetByEmailFunc) PushReturn(r0 *types.AccessRequest, r1 error) {
	f.PushHook(func(context.Context, string) (*types.AccessRequest, error) {
		return r0, r1
	})
}

func (f *AccessRequestStoreGetByEmailFunc) nextHook() func(context.Context, string) (*types.AccessRequest, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *AccessRequestStoreGetByEmailFunc) appendCall(r0 AccessRequestStoreGetByEmailFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of AccessRequestStoreGetByEmailFuncCall
// objects describing the invocations of this function.
func (f *AccessRequestStoreGetByEmailFunc) History() []AccessRequestStoreGetByEmailFuncCall {
	f.mutex.Lock()
	history := make([]AccessRequestStoreGetByEmailFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// AccessRequestStoreGetByEmailFuncCall is an object that describes an
// invocation of method GetByEmail on an instance of MockAccessRequestStore.
type AccessRequestStoreGetByEmailFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *types.AccessRequest
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c AccessRequestStoreGetByEmailFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c AccessRequestStoreGetByEmailFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// AccessRequestStoreGetByIDFunc describes the behavior when the GetByID
// method of the parent MockAccessRequestStore instance is invoked.
type AccessRequestStoreGetByIDFunc struct {
	defaultHook func(context.Context, int32) (*types.AccessRequest, error)
	hooks       []func(context.Context, int32) (*types.AccessRequest, error)
	history     []AccessRequestStoreGetByIDFuncCall
	mutex       sync.Mutex
}

// GetByID delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockAccessRequestStore) GetByID(v0 context.Context, v1 int32) (*types.AccessRequest, error) {
	r0, r1 := m.GetByIDFunc.nextHook()(v0, v1)
	m.GetByIDFunc.appendCall(AccessRequestStoreGetByIDFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetByID method of
// the parent MockAccessRequestStore instance is invoked and the hook queue
// is empty.
func (f *AccessRequestStoreGetByIDFunc) SetDefaultHook(hook func(context.Context, int32) (*types.AccessRequest, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetByID method of the parent MockAccessRequestStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *AccessRequestStoreGetByIDFunc) PushHook(hook func(context.Context, int32) (*types.AccessRequest, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *AccessRequestStoreGetByIDFunc) SetDefaultReturn(r0 *types.AccessRequest, r1 error) {
	f.SetDefaultHook(func(context.Context, int32) (*types.AccessRequest, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *AccessRequestStoreGetByIDFunc) PushReturn(r0 *types.AccessRequest, r1 error) {
	f.PushHook(func(context.Context, int32) (*types.AccessRequest, error) {
		return r0, r1
	})
}

func (f *AccessRequestStoreGetByIDFunc) nextHook() func(context.Context, int32) (*types.AccessRequest, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *AccessRequestStoreGetByIDFunc) appendCall(r0 AccessRequestStoreGetByIDFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of AccessRequestStoreGetByIDFuncCall objects
// describing the invocations of this function.
func (f *AccessRequestStoreGetByIDFunc) History() []AccessRequestStoreGetByIDFuncCall {
	f.mutex.Lock()
	history := make([]AccessRequestStoreGetByIDFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// AccessRequestStoreGetByIDFuncCall is an object that describes an
// invocation of method GetByID on an instance of MockAccessRequestStore.
type AccessRequestStoreGetByIDFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int32
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *types.AccessRequest
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c AccessRequestStoreGetByIDFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c AccessRequestStoreGetByIDFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// AccessRequestStoreHandleFunc describes the behavior when the Handle
// method of the parent MockAccessRequestStore instance is invoked.
type AccessRequestStoreHandleFunc struct {
	defaultHook func() basestore.TransactableHandle
	hooks       []func() basestore.TransactableHandle
	history     []AccessRequestStoreHandleFuncCall
	mutex       sync.Mutex
}

// Handle delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockAccessRequestStore) Handle() basestore.TransactableHandle {
	r0 := m.HandleFunc.nextHook()()
	m.HandleFunc.appendCall(AccessRequestStoreHandleFuncCall{r0})
	return r0
}

// SetDefaultHook sets function that is called when the Handle method of the
// parent MockAccessRequestStore instance is invoked and the hook queue is
// empty.
func (f *AccessRequestStoreHandleFunc) SetDefaultHook(hook func() basestore.TransactableHandle) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Handle method of the parent MockAccessRequestStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *AccessRequestStoreHandleFunc) PushHook(hook func() basestore.TransactableHandle) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *AccessRequestStoreHandleFunc) SetDefaultReturn(r0 basestore.TransactableHandle) {
	f.SetDefaultHook(func() basestore.TransactableHandle {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *AccessRequestStoreHandleFunc) PushReturn(r0 basestore.TransactableHandle) {
	f.PushHook(func() basestore.TransactableHandle {
		return r0
	})
}

func (f *AccessRequestStoreHandleFunc) nextHook() func() basestore.TransactableHandle {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *AccessRequestStoreHandleFunc) appendCall(r0 AccessRequestStoreHandleFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of AccessRequestStoreHandleFuncCall objects
// describing the invocations of this function.
func (f *AccessRequestStoreHandleFunc) History() []AccessRequestStoreHandleFuncCall {
	f.mutex.Lock()
	history := make([]AccessRequestStoreHandleFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// AccessRequestStoreHandleFuncCall is an object that describes an
// invocation of method Handle on an instance of MockAccessRequestStore.
type AccessRequestStoreHandleFuncCall struct {
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 basestore.TransactableHandle
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c AccessRequestStoreHandleFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c AccessRequestStoreHandleFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// AccessRequestStoreListFunc describes the behavior when the List method of
// the parent MockAccessRequestStore instance is invoked.
type AccessRequestStoreListFunc struct {
	defaultHook func(context.Context, *AccessRequestsFilterArgs, *database.PaginationArgs) ([]*types.AccessRequest, error)
	hooks       []func(context.Context, *AccessRequestsFilterArgs, *database.PaginationArgs) ([]*types.AccessRequest, error)
	history     []AccessRequestStoreListFuncCall
	mutex       sync.Mutex
}

// List delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockAccessRequestStore) List(v0 context.Context, v1 *AccessRequestsFilterArgs, v2 *database.PaginationArgs) ([]*types.AccessRequest, error) {
	r0, r1 := m.ListFunc.nextHook()(v0, v1, v2)
	m.ListFunc.appendCall(AccessRequestStoreListFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the List method of the
// parent MockAccessRequestStore instance is invoked and the hook queue is
// empty.
func (f *AccessRequestStoreListFunc) SetDefaultHook(hook func(context.Context, *AccessRequestsFilterArgs, *database.PaginationArgs) ([]*types.AccessRequest, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// List method of the parent MockAccessRequestStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *AccessRequestStoreListFunc) PushHook(hook func(context.Context, *AccessRequestsFilterArgs, *database.PaginationArgs) ([]*types.AccessRequest, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *AccessRequestStoreListFunc) SetDefaultReturn(r0 []*types.AccessRequest, r1 error) {
	f.SetDefaultHook(func(context.Context, *AccessRequestsFilterArgs, *database.PaginationArgs) ([]*types.AccessRequest, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *AccessRequestStoreListFunc) PushReturn(r0 []*types.AccessRequest, r1 error) {
	f.PushHook(func(context.Context, *AccessRequestsFilterArgs, *database.PaginationArgs) ([]*types.AccessRequest, error) {
		return r0, r1
	})
}

func (f *AccessRequestStoreListFunc) nextHook() func(context.Context, *AccessRequestsFilterArgs, *database.PaginationArgs) ([]*types.AccessRequest, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *AccessRequestStoreListFunc) appendCall(r0 AccessRequestStoreListFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of AccessRequestStoreListFuncCall objects
// describing the invocations of this function.
func (f *AccessRequestStoreListFunc) History() []AccessRequestStoreListFuncCall {
	f.mutex.Lock()
	history := make([]AccessRequestStoreListFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// AccessRequestStoreListFuncCall is an object that describes an invocation
// of method List on an instance of MockAccessRequestStore.
type AccessRequestStoreListFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 *AccessRequestsFilterArgs
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 *database.PaginationArgs
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []*types.AccessRequest
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c AccessRequestStoreListFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c AccessRequestStoreListFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// AccessRequestStoreUpdateFunc describes the behavior when the Update
// method of the parent MockAccessRequestStore instance is invoked.
type AccessRequestStoreUpdateFunc struct {
	defaultHook func(context.Context, *types.AccessRequest) (*types.AccessRequest, error)
	hooks       []func(context.Context, *types.AccessRequest) (*types.AccessRequest, error)
	history     []AccessRequestStoreUpdateFuncCall
	mutex       sync.Mutex
}

// Update delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockAccessRequestStore) Update(v0 context.Context, v1 *types.AccessRequest) (*types.AccessRequest, error) {
	r0, r1 := m.UpdateFunc.nextHook()(v0, v1)
	m.UpdateFunc.appendCall(AccessRequestStoreUpdateFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the Update method of the
// parent MockAccessRequestStore instance is invoked and the hook queue is
// empty.
func (f *AccessRequestStoreUpdateFunc) SetDefaultHook(hook func(context.Context, *types.AccessRequest) (*types.AccessRequest, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Update method of the parent MockAccessRequestStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *AccessRequestStoreUpdateFunc) PushHook(hook func(context.Context, *types.AccessRequest) (*types.AccessRequest, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *AccessRequestStoreUpdateFunc) SetDefaultReturn(r0 *types.AccessRequest, r1 error) {
	f.SetDefaultHook(func(context.Context, *types.AccessRequest) (*types.AccessRequest, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *AccessRequestStoreUpdateFunc) PushReturn(r0 *types.AccessRequest, r1 error) {
	f.PushHook(func(context.Context, *types.AccessRequest) (*types.AccessRequest, error) {
		return r0, r1
	})
}

func (f *AccessRequestStoreUpdateFunc) nextHook() func(context.Context, *types.AccessRequest) (*types.AccessRequest, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *AccessRequestStoreUpdateFunc) appendCall(r0 AccessRequestStoreUpdateFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of AccessRequestStoreUpdateFuncCall objects
// describing the invocations of this function.
func (f *AccessRequestStoreUpdateFunc) History() []AccessRequestStoreUpdateFuncCall {
	f.mutex.Lock()
	history := make([]AccessRequestStoreUpdateFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// AccessRequestStoreUpdateFuncCall is an object that describes an
// invocation of method Update on an instance of MockAccessRequestStore.
type AccessRequestStoreUpdateFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 *types.AccessRequest
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *types.AccessRequest
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c AccessRequestStoreUpdateFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c AccessRequestStoreUpdateFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// AccessRequestStoreWithTransactFunc describes the behavior when the
// WithTransact method of the parent MockAccessRequestStore instance is
// invoked.
type AccessRequestStoreWithTransactFunc struct {
	defaultHook func(context.Context, func(AccessRequestStore) error) error
	hooks       []func(context.Context, func(AccessRequestStore) error) error
	history     []AccessRequestStoreWithTransactFuncCall
	mutex       sync.Mutex
}

// WithTransact delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockAccessRequestStore) WithTransact(v0 context.Context, v1 func(AccessRequestStore) error) error {
	r0 := m.WithTransactFunc.nextHook()(v0, v1)
	m.WithTransactFunc.appendCall(AccessRequestStoreWithTransactFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the WithTransact method
// of the parent MockAccessRequestStore instance is invoked and the hook
// queue is empty.
func (f *AccessRequestStoreWithTransactFunc) SetDefaultHook(hook func(context.Context, func(AccessRequestStore) error) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// WithTransact method of the parent MockAccessRequestStore instance invokes
// the hook at the front of the queue and discards it. After the queue is
// empty, the default hook function is invoked for any future action.
func (f *AccessRequestStoreWithTransactFunc) PushHook(hook func(context.Context, func(AccessRequestStore) error) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *AccessRequestStoreWithTransactFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, func(AccessRequestStore) error) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *AccessRequestStoreWithTransactFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, func(AccessRequestStore) error) error {
		return r0
	})
}

func (f *AccessRequestStoreWithTransactFunc) nextHook() func(context.Context, func(AccessRequestStore) error) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *AccessRequestStoreWithTransactFunc) appendCall(r0 AccessRequestStoreWithTransactFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of AccessRequestStoreWithTransactFuncCall
// objects describing the invocations of this function.
func (f *AccessRequestStoreWithTransactFunc) History() []AccessRequestStoreWithTransactFuncCall {
	f.mutex.Lock()
	history := make([]AccessRequestStoreWithTransactFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// AccessRequestStoreWithTransactFuncCall is an object that describes an
// invocation of method WithTransact on an instance of
// MockAccessRequestStore.
type AccessRequestStoreWithTransactFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 func(AccessRequestStore) error
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c AccessRequestStoreWithTransactFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c AccessRequestStoreWithTransactFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}
