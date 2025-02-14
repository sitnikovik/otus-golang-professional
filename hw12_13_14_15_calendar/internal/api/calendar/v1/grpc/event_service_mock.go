// Code generated by mockery v2.43.2. DO NOT EDIT.

package grpc

import (
	context "context"

	filterevent "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"
	event "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"

	mock "github.com/stretchr/testify/mock"
)

// MockeventService is an autogenerated mock type for the eventService type
type MockeventService struct {
	mock.Mock
}

type MockeventService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockeventService) EXPECT() *MockeventService_Expecter {
	return &MockeventService_Expecter{mock: &_m.Mock}
}

// CreateEvent provides a mock function with given fields: ctx, _a1
func (_m *MockeventService) CreateEvent(ctx context.Context, _a1 *event.Event) (uint64, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateEvent")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *event.Event) (uint64, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *event.Event) uint64); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *event.Event) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockeventService_CreateEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateEvent'
type MockeventService_CreateEvent_Call struct {
	*mock.Call
}

// CreateEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *event.Event
func (_e *MockeventService_Expecter) CreateEvent(ctx interface{}, _a1 interface{}) *MockeventService_CreateEvent_Call {
	return &MockeventService_CreateEvent_Call{Call: _e.mock.On("CreateEvent", ctx, _a1)}
}

func (_c *MockeventService_CreateEvent_Call) Run(run func(ctx context.Context, _a1 *event.Event)) *MockeventService_CreateEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*event.Event))
	})
	return _c
}

func (_c *MockeventService_CreateEvent_Call) Return(_a0 uint64, _a1 error) *MockeventService_CreateEvent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockeventService_CreateEvent_Call) RunAndReturn(run func(context.Context, *event.Event) (uint64, error)) *MockeventService_CreateEvent_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteEvent provides a mock function with given fields: ctx, eventID
func (_m *MockeventService) DeleteEvent(ctx context.Context, eventID uint64) error {
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

// MockeventService_DeleteEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteEvent'
type MockeventService_DeleteEvent_Call struct {
	*mock.Call
}

// DeleteEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - eventID uint64
func (_e *MockeventService_Expecter) DeleteEvent(ctx interface{}, eventID interface{}) *MockeventService_DeleteEvent_Call {
	return &MockeventService_DeleteEvent_Call{Call: _e.mock.On("DeleteEvent", ctx, eventID)}
}

func (_c *MockeventService_DeleteEvent_Call) Run(run func(ctx context.Context, eventID uint64)) *MockeventService_DeleteEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *MockeventService_DeleteEvent_Call) Return(_a0 error) *MockeventService_DeleteEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockeventService_DeleteEvent_Call) RunAndReturn(run func(context.Context, uint64) error) *MockeventService_DeleteEvent_Call {
	_c.Call.Return(run)
	return _c
}

// GetEvent provides a mock function with given fields: ctx, eventID
func (_m *MockeventService) GetEvent(ctx context.Context, eventID uint64) (*event.Event, error) {
	ret := _m.Called(ctx, eventID)

	if len(ret) == 0 {
		panic("no return value specified for GetEvent")
	}

	var r0 *event.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*event.Event, error)); ok {
		return rf(ctx, eventID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *event.Event); ok {
		r0 = rf(ctx, eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*event.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, eventID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockeventService_GetEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEvent'
type MockeventService_GetEvent_Call struct {
	*mock.Call
}

// GetEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - eventID uint64
func (_e *MockeventService_Expecter) GetEvent(ctx interface{}, eventID interface{}) *MockeventService_GetEvent_Call {
	return &MockeventService_GetEvent_Call{Call: _e.mock.On("GetEvent", ctx, eventID)}
}

func (_c *MockeventService_GetEvent_Call) Run(run func(ctx context.Context, eventID uint64)) *MockeventService_GetEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *MockeventService_GetEvent_Call) Return(_a0 *event.Event, _a1 error) *MockeventService_GetEvent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockeventService_GetEvent_Call) RunAndReturn(run func(context.Context, uint64) (*event.Event, error)) *MockeventService_GetEvent_Call {
	_c.Call.Return(run)
	return _c
}

// GetEvents provides a mock function with given fields: ctx, filter
func (_m *MockeventService) GetEvents(ctx context.Context, filter filterevent.Filter) ([]*event.Event, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetEvents")
	}

	var r0 []*event.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, filterevent.Filter) ([]*event.Event, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, filterevent.Filter) []*event.Event); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*event.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, filterevent.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockeventService_GetEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEvents'
type MockeventService_GetEvents_Call struct {
	*mock.Call
}

// GetEvents is a helper method to define mock.On call
//   - ctx context.Context
//   - filter filterevent.Filter
func (_e *MockeventService_Expecter) GetEvents(ctx interface{}, filter interface{}) *MockeventService_GetEvents_Call {
	return &MockeventService_GetEvents_Call{Call: _e.mock.On("GetEvents", ctx, filter)}
}

func (_c *MockeventService_GetEvents_Call) Run(run func(ctx context.Context, filter filterevent.Filter)) *MockeventService_GetEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(filterevent.Filter))
	})
	return _c
}

func (_c *MockeventService_GetEvents_Call) Return(_a0 []*event.Event, _a1 error) *MockeventService_GetEvents_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockeventService_GetEvents_Call) RunAndReturn(run func(context.Context, filterevent.Filter) ([]*event.Event, error)) *MockeventService_GetEvents_Call {
	_c.Call.Return(run)
	return _c
}

// GetEventsBeforeDays provides a mock function with given fields: ctx, days
func (_m *MockeventService) GetEventsBeforeDays(ctx context.Context, days uint32) ([]*event.Event, error) {
	ret := _m.Called(ctx, days)

	if len(ret) == 0 {
		panic("no return value specified for GetEventsBeforeDays")
	}

	var r0 []*event.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) ([]*event.Event, error)); ok {
		return rf(ctx, days)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) []*event.Event); ok {
		r0 = rf(ctx, days)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*event.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, days)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockeventService_GetEventsBeforeDays_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEventsBeforeDays'
type MockeventService_GetEventsBeforeDays_Call struct {
	*mock.Call
}

// GetEventsBeforeDays is a helper method to define mock.On call
//   - ctx context.Context
//   - days uint32
func (_e *MockeventService_Expecter) GetEventsBeforeDays(ctx interface{}, days interface{}) *MockeventService_GetEventsBeforeDays_Call {
	return &MockeventService_GetEventsBeforeDays_Call{Call: _e.mock.On("GetEventsBeforeDays", ctx, days)}
}

func (_c *MockeventService_GetEventsBeforeDays_Call) Run(run func(ctx context.Context, days uint32)) *MockeventService_GetEventsBeforeDays_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *MockeventService_GetEventsBeforeDays_Call) Return(_a0 []*event.Event, _a1 error) *MockeventService_GetEventsBeforeDays_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockeventService_GetEventsBeforeDays_Call) RunAndReturn(run func(context.Context, uint32) ([]*event.Event, error)) *MockeventService_GetEventsBeforeDays_Call {
	_c.Call.Return(run)
	return _c
}

// GetEventsForMonth provides a mock function with given fields: ctx
func (_m *MockeventService) GetEventsForMonth(ctx context.Context) ([]*event.Event, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetEventsForMonth")
	}

	var r0 []*event.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*event.Event, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*event.Event); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*event.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockeventService_GetEventsForMonth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEventsForMonth'
type MockeventService_GetEventsForMonth_Call struct {
	*mock.Call
}

// GetEventsForMonth is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockeventService_Expecter) GetEventsForMonth(ctx interface{}) *MockeventService_GetEventsForMonth_Call {
	return &MockeventService_GetEventsForMonth_Call{Call: _e.mock.On("GetEventsForMonth", ctx)}
}

func (_c *MockeventService_GetEventsForMonth_Call) Run(run func(ctx context.Context)) *MockeventService_GetEventsForMonth_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockeventService_GetEventsForMonth_Call) Return(_a0 []*event.Event, _a1 error) *MockeventService_GetEventsForMonth_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockeventService_GetEventsForMonth_Call) RunAndReturn(run func(context.Context) ([]*event.Event, error)) *MockeventService_GetEventsForMonth_Call {
	_c.Call.Return(run)
	return _c
}

// GetEventsForToday provides a mock function with given fields: ctx
func (_m *MockeventService) GetEventsForToday(ctx context.Context) ([]*event.Event, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetEventsForToday")
	}

	var r0 []*event.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*event.Event, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*event.Event); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*event.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockeventService_GetEventsForToday_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEventsForToday'
type MockeventService_GetEventsForToday_Call struct {
	*mock.Call
}

// GetEventsForToday is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockeventService_Expecter) GetEventsForToday(ctx interface{}) *MockeventService_GetEventsForToday_Call {
	return &MockeventService_GetEventsForToday_Call{Call: _e.mock.On("GetEventsForToday", ctx)}
}

func (_c *MockeventService_GetEventsForToday_Call) Run(run func(ctx context.Context)) *MockeventService_GetEventsForToday_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockeventService_GetEventsForToday_Call) Return(_a0 []*event.Event, _a1 error) *MockeventService_GetEventsForToday_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockeventService_GetEventsForToday_Call) RunAndReturn(run func(context.Context) ([]*event.Event, error)) *MockeventService_GetEventsForToday_Call {
	_c.Call.Return(run)
	return _c
}

// GetEventsForWeek provides a mock function with given fields: ctx
func (_m *MockeventService) GetEventsForWeek(ctx context.Context) ([]*event.Event, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetEventsForWeek")
	}

	var r0 []*event.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*event.Event, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*event.Event); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*event.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockeventService_GetEventsForWeek_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEventsForWeek'
type MockeventService_GetEventsForWeek_Call struct {
	*mock.Call
}

// GetEventsForWeek is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockeventService_Expecter) GetEventsForWeek(ctx interface{}) *MockeventService_GetEventsForWeek_Call {
	return &MockeventService_GetEventsForWeek_Call{Call: _e.mock.On("GetEventsForWeek", ctx)}
}

func (_c *MockeventService_GetEventsForWeek_Call) Run(run func(ctx context.Context)) *MockeventService_GetEventsForWeek_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockeventService_GetEventsForWeek_Call) Return(_a0 []*event.Event, _a1 error) *MockeventService_GetEventsForWeek_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockeventService_GetEventsForWeek_Call) RunAndReturn(run func(context.Context) ([]*event.Event, error)) *MockeventService_GetEventsForWeek_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateEvent provides a mock function with given fields: ctx, _a1
func (_m *MockeventService) UpdateEvent(ctx context.Context, _a1 *event.Event) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *event.Event) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockeventService_UpdateEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateEvent'
type MockeventService_UpdateEvent_Call struct {
	*mock.Call
}

// UpdateEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *event.Event
func (_e *MockeventService_Expecter) UpdateEvent(ctx interface{}, _a1 interface{}) *MockeventService_UpdateEvent_Call {
	return &MockeventService_UpdateEvent_Call{Call: _e.mock.On("UpdateEvent", ctx, _a1)}
}

func (_c *MockeventService_UpdateEvent_Call) Run(run func(ctx context.Context, _a1 *event.Event)) *MockeventService_UpdateEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*event.Event))
	})
	return _c
}

func (_c *MockeventService_UpdateEvent_Call) Return(_a0 error) *MockeventService_UpdateEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockeventService_UpdateEvent_Call) RunAndReturn(run func(context.Context, *event.Event) error) *MockeventService_UpdateEvent_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockeventService creates a new instance of MockeventService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockeventService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockeventService {
	mock := &MockeventService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
