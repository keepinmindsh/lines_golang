// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Tire is an autogenerated mock type for the Tire type
type Tire struct {
	mock.Mock
}

type Tire_Expecter struct {
	mock *mock.Mock
}

func (_m *Tire) EXPECT() *Tire_Expecter {
	return &Tire_Expecter{mock: &_m.Mock}
}

// Backward provides a mock function with given fields:
func (_m *Tire) Backward() {
	_m.Called()
}

// Tire_Backward_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Backward'
type Tire_Backward_Call struct {
	*mock.Call
}

// Backward is a helper method to define mock.On call
func (_e *Tire_Expecter) Backward() *Tire_Backward_Call {
	return &Tire_Backward_Call{Call: _e.mock.On("Backward")}
}

func (_c *Tire_Backward_Call) Run(run func()) *Tire_Backward_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tire_Backward_Call) Return() *Tire_Backward_Call {
	_c.Call.Return()
	return _c
}

func (_c *Tire_Backward_Call) RunAndReturn(run func()) *Tire_Backward_Call {
	_c.Call.Return(run)
	return _c
}

// Forward provides a mock function with given fields:
func (_m *Tire) Forward() {
	_m.Called()
}

// Tire_Forward_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Forward'
type Tire_Forward_Call struct {
	*mock.Call
}

// Forward is a helper method to define mock.On call
func (_e *Tire_Expecter) Forward() *Tire_Forward_Call {
	return &Tire_Forward_Call{Call: _e.mock.On("Forward")}
}

func (_c *Tire_Forward_Call) Run(run func()) *Tire_Forward_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tire_Forward_Call) Return() *Tire_Forward_Call {
	_c.Call.Return()
	return _c
}

func (_c *Tire_Forward_Call) RunAndReturn(run func()) *Tire_Forward_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields:
func (_m *Tire) Start() {
	_m.Called()
}

// Tire_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Tire_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *Tire_Expecter) Start() *Tire_Start_Call {
	return &Tire_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *Tire_Start_Call) Run(run func()) *Tire_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tire_Start_Call) Return() *Tire_Start_Call {
	_c.Call.Return()
	return _c
}

func (_c *Tire_Start_Call) RunAndReturn(run func()) *Tire_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *Tire) Stop() {
	_m.Called()
}

// Tire_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type Tire_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *Tire_Expecter) Stop() *Tire_Stop_Call {
	return &Tire_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *Tire_Stop_Call) Run(run func()) *Tire_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tire_Stop_Call) Return() *Tire_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *Tire_Stop_Call) RunAndReturn(run func()) *Tire_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// NewTire creates a new instance of Tire. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTire(t interface {
	mock.TestingT
	Cleanup(func())
}) *Tire {
	mock := &Tire{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
