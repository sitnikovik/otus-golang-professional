// Code generated by mockery v2.43.2. DO NOT EDIT.

package event

import (
	context "context"

	filterevent "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"

	mock "github.com/stretchr/testify/mock"

	modelevent "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

// MockeventDB is an autogenerated mock type for the eventDB type
type MockeventDB struct {
	mock.Mock
}

type MockeventDB_Expecter struct {
	mock *mock.Mock
}

func (_m *MockeventDB) EXPECT() *MockeventDB_Expecter {
	return &MockeventDB_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields: ctx
func (_m *MockeventDB) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockeventDB_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockeventDB_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockeventDB_Expecter) Close(ctx interface{}) *MockeventDB_Close_Call {
	return &MockeventDB_Close_Call{Call: _e.mock.On("Close", ctx)}
}

func (_c *MockeventDB_Close_Call) Run(run func(ctx context.Context)) *MockeventDB_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockeventDB_Close_Call) Return(_a0 error) *MockeventDB_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockeventDB_Close_Call) RunAndReturn(run func(context.Context) error) *MockeventDB_Close_Call {
	_c.Call.Return(run)
	return _c
}

// CreateEvent provides a mock function with given fields: ctx, _a1
func (_m *MockeventDB) CreateEvent(ctx context.Context, _a1 *modelevent.Event) (uint64, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateEvent")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *modelevent.Event) (uint64, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *modelevent.Event) uint64); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *modelevent.Event) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockeventDB_CreateEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateEvent'
type MockeventDB_CreateEvent_Call struct {
	*mock.Call
}

// CreateEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *modelevent.Event
func (_e *MockeventDB_Expecter) CreateEvent(ctx interface{}, _a1 interface{}) *MockeventDB_CreateEvent_Call {
	return &MockeventDB_CreateEvent_Call{Call: _e.mock.On("CreateEvent", ctx, _a1)}
}

func (_c *MockeventDB_CreateEvent_Call) Run(run func(ctx context.Context, _a1 *modelevent.Event)) *MockeventDB_CreateEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*modelevent.Event))
	})
	return _c
}

func (_c *MockeventDB_CreateEvent_Call) Return(_a0 uint64, _a1 error) *MockeventDB_CreateEvent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockeventDB_CreateEvent_Call) RunAndReturn(run func(context.Context, *modelevent.Event) (uint64, error)) *MockeventDB_CreateEvent_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteEvent provides a mock function with given fields: ctx, eventID
func (_m *MockeventDB) DeleteEvent(ctx context.Context, eventID uint64) error {
	ret := _m.Called(ctx, eventID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, eventID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockeventDB_DeleteEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteEvent'
type MockeventDB_DeleteEvent_Call struct {
	*mock.Call
}

// DeleteEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - eventID uint64
func (_e *MockeventDB_Expecter) DeleteEvent(ctx interface{}, eventID interface{}) *MockeventDB_DeleteEvent_Call {
	return &MockeventDB_DeleteEvent_Call{Call: _e.mock.On("DeleteEvent", ctx, eventID)}
}

func (_c *MockeventDB_DeleteEvent_Call) Run(run func(ctx context.Context, eventID uint64)) *MockeventDB_DeleteEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *MockeventDB_DeleteEvent_Call) Return(_a0 error) *MockeventDB_DeleteEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockeventDB_DeleteEvent_Call) RunAndReturn(run func(context.Context, uint64) error) *MockeventDB_DeleteEvent_Call {
	_c.Call.Return(run)
	return _c
}

// GetEvent provides a mock function with given fields: ctx, eventID
func (_m *MockeventDB) GetEvent(ctx context.Context, eventID uint64) (*modelevent.Event, error) {
	ret := _m.Called(ctx, eventID)

	if len(ret) == 0 {
		panic("no return value specified for GetEvent")
	}

	var r0 *modelevent.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*modelevent.Event, error)); ok {
		return rf(ctx, eventID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *modelevent.Event); ok {
		r0 = rf(ctx, eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*modelevent.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, eventID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockeventDB_GetEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEvent'
type MockeventDB_GetEvent_Call struct {
	*mock.Call
}

// GetEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - eventID uint64
func (_e *MockeventDB_Expecter) GetEvent(ctx interface{}, eventID interface{}) *MockeventDB_GetEvent_Call {
	return &MockeventDB_GetEvent_Call{Call: _e.mock.On("GetEvent", ctx, eventID)}
}

func (_c *MockeventDB_GetEvent_Call) Run(run func(ctx context.Context, eventID uint64)) *MockeventDB_GetEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *MockeventDB_GetEvent_Call) Return(_a0 *modelevent.Event, _a1 error) *MockeventDB_GetEvent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockeventDB_GetEvent_Call) RunAndReturn(run func(context.Context, uint64) (*modelevent.Event, error)) *MockeventDB_GetEvent_Call {
	_c.Call.Return(run)
	return _c
}

// GetEvents provides a mock function with given fields: ctx, filter
func (_m *MockeventDB) GetEvents(ctx context.Context, filter filterevent.Filter) ([]*modelevent.Event, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetEvents")
	}

	var r0 []*modelevent.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, filterevent.Filter) ([]*modelevent.Event, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, filterevent.Filter) []*modelevent.Event); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*modelevent.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, filterevent.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockeventDB_GetEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEvents'
type MockeventDB_GetEvents_Call struct {
	*mock.Call
}

// GetEvents is a helper method to define mock.On call
//   - ctx context.Context
//   - filter filterevent.Filter
func (_e *MockeventDB_Expecter) GetEvents(ctx interface{}, filter interface{}) *MockeventDB_GetEvents_Call {
	return &MockeventDB_GetEvents_Call{Call: _e.mock.On("GetEvents", ctx, filter)}
}

func (_c *MockeventDB_GetEvents_Call) Run(run func(ctx context.Context, filter filterevent.Filter)) *MockeventDB_GetEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(filterevent.Filter))
	})
	return _c
}

func (_c *MockeventDB_GetEvents_Call) Return(_a0 []*modelevent.Event, _a1 error) *MockeventDB_GetEvents_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockeventDB_GetEvents_Call) RunAndReturn(run func(context.Context, filterevent.Filter) ([]*modelevent.Event, error)) *MockeventDB_GetEvents_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateEvent provides a mock function with given fields: ctx, _a1
func (_m *MockeventDB) UpdateEvent(ctx context.Context, _a1 *modelevent.Event) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *modelevent.Event) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockeventDB_UpdateEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateEvent'
type MockeventDB_UpdateEvent_Call struct {
	*mock.Call
}

// UpdateEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *modelevent.Event
func (_e *MockeventDB_Expecter) UpdateEvent(ctx interface{}, _a1 interface{}) *MockeventDB_UpdateEvent_Call {
	return &MockeventDB_UpdateEvent_Call{Call: _e.mock.On("UpdateEvent", ctx, _a1)}
}

func (_c *MockeventDB_UpdateEvent_Call) Run(run func(ctx context.Context, _a1 *modelevent.Event)) *MockeventDB_UpdateEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*modelevent.Event))
	})
	return _c
}

func (_c *MockeventDB_UpdateEvent_Call) Return(_a0 error) *MockeventDB_UpdateEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockeventDB_UpdateEvent_Call) RunAndReturn(run func(context.Context, *modelevent.Event) error) *MockeventDB_UpdateEvent_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockeventDB creates a new instance of MockeventDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockeventDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockeventDB {
	mock := &MockeventDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
