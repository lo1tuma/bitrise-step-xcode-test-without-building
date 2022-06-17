// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Xcodebuild is an autogenerated mock type for the Xcodebuild type
type Xcodebuild struct {
	mock.Mock
}

// TestWithoutBuilding provides a mock function with given fields: xctestrun, destination, testRepetitionMode, maximumTestRepetitions, relaunchTestsForEachRepetition, options
func (_m *Xcodebuild) TestWithoutBuilding(xctestrun string, destination string, testRepetitionMode string, maximumTestRepetitions int, relaunchTestsForEachRepetition bool, options ...string) (string, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, xctestrun, destination, testRepetitionMode, maximumTestRepetitions, relaunchTestsForEachRepetition)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, string, int, bool, ...string) string); ok {
		r0 = rf(xctestrun, destination, testRepetitionMode, maximumTestRepetitions, relaunchTestsForEachRepetition, options...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, int, bool, ...string) error); ok {
		r1 = rf(xctestrun, destination, testRepetitionMode, maximumTestRepetitions, relaunchTestsForEachRepetition, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}