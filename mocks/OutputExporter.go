// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// OutputExporter is an autogenerated mock type for the OutputExporter type
type OutputExporter struct {
	mock.Mock
}

// CopyAndSaveTestData provides a mock function with given fields: sourceTestOutputDir, targetAddonPath, targetAddonBundleName
func (_m *OutputExporter) CopyAndSaveTestData(sourceTestOutputDir string, targetAddonPath string, targetAddonBundleName string) error {
	ret := _m.Called(sourceTestOutputDir, targetAddonPath, targetAddonBundleName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(sourceTestOutputDir, targetAddonPath, targetAddonBundleName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ZipAndExportOutput provides a mock function with given fields: artifact, destinationZipPth, envKey
func (_m *OutputExporter) ZipAndExportOutput(artifact string, destinationZipPth string, envKey string) error {
	ret := _m.Called(artifact, destinationZipPth, envKey)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(artifact, destinationZipPth, envKey)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
