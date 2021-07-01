// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	scanners "github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client/scanners"
)

// MockScannersClientService is an autogenerated mock type for the ClientService type
type MockScannersClientService struct {
	mock.Mock
}

// DeleteScannersRegistrationID provides a mock function with given fields: params, authInfo, opts
func (_m *MockScannersClientService) DeleteScannersRegistrationID(params *scanners.DeleteScannersRegistrationIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...scanners.ClientOption) (*scanners.DeleteScannersRegistrationIDOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *scanners.DeleteScannersRegistrationIDOK
	if rf, ok := ret.Get(0).(func(*scanners.DeleteScannersRegistrationIDParams, runtime.ClientAuthInfoWriter, ...scanners.ClientOption) *scanners.DeleteScannersRegistrationIDOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scanners.DeleteScannersRegistrationIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scanners.DeleteScannersRegistrationIDParams, runtime.ClientAuthInfoWriter, ...scanners.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PatchScannersRegistrationID provides a mock function with given fields: params, authInfo, opts
func (_m *MockScannersClientService) PatchScannersRegistrationID(params *scanners.PatchScannersRegistrationIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...scanners.ClientOption) (*scanners.PatchScannersRegistrationIDOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *scanners.PatchScannersRegistrationIDOK
	if rf, ok := ret.Get(0).(func(*scanners.PatchScannersRegistrationIDParams, runtime.ClientAuthInfoWriter, ...scanners.ClientOption) *scanners.PatchScannersRegistrationIDOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scanners.PatchScannersRegistrationIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scanners.PatchScannersRegistrationIDParams, runtime.ClientAuthInfoWriter, ...scanners.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostScanners provides a mock function with given fields: params, authInfo, opts
func (_m *MockScannersClientService) PostScanners(params *scanners.PostScannersParams, authInfo runtime.ClientAuthInfoWriter, opts ...scanners.ClientOption) (*scanners.PostScannersCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *scanners.PostScannersCreated
	if rf, ok := ret.Get(0).(func(*scanners.PostScannersParams, runtime.ClientAuthInfoWriter, ...scanners.ClientOption) *scanners.PostScannersCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scanners.PostScannersCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scanners.PostScannersParams, runtime.ClientAuthInfoWriter, ...scanners.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutProjectsProjectIDScanner provides a mock function with given fields: params, authInfo, opts
func (_m *MockScannersClientService) PutProjectsProjectIDScanner(params *scanners.PutProjectsProjectIDScannerParams, authInfo runtime.ClientAuthInfoWriter, opts ...scanners.ClientOption) (*scanners.PutProjectsProjectIDScannerOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *scanners.PutProjectsProjectIDScannerOK
	if rf, ok := ret.Get(0).(func(*scanners.PutProjectsProjectIDScannerParams, runtime.ClientAuthInfoWriter, ...scanners.ClientOption) *scanners.PutProjectsProjectIDScannerOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scanners.PutProjectsProjectIDScannerOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scanners.PutProjectsProjectIDScannerParams, runtime.ClientAuthInfoWriter, ...scanners.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutScannersRegistrationID provides a mock function with given fields: params, authInfo, opts
func (_m *MockScannersClientService) PutScannersRegistrationID(params *scanners.PutScannersRegistrationIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...scanners.ClientOption) (*scanners.PutScannersRegistrationIDOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *scanners.PutScannersRegistrationIDOK
	if rf, ok := ret.Get(0).(func(*scanners.PutScannersRegistrationIDParams, runtime.ClientAuthInfoWriter, ...scanners.ClientOption) *scanners.PutScannersRegistrationIDOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scanners.PutScannersRegistrationIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scanners.PutScannersRegistrationIDParams, runtime.ClientAuthInfoWriter, ...scanners.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockScannersClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}
