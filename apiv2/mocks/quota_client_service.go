// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	quota "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/quota"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// MockQuotaClientService is an autogenerated mock type for the ClientService type
type MockQuotaClientService struct {
	mock.Mock
}

// GetQuota provides a mock function with given fields: params, authInfo
func (_m *MockQuotaClientService) GetQuota(params *quota.GetQuotaParams, authInfo runtime.ClientAuthInfoWriter) (*quota.GetQuotaOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *quota.GetQuotaOK
	if rf, ok := ret.Get(0).(func(*quota.GetQuotaParams, runtime.ClientAuthInfoWriter) *quota.GetQuotaOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*quota.GetQuotaOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*quota.GetQuotaParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListQuotas provides a mock function with given fields: params, authInfo
func (_m *MockQuotaClientService) ListQuotas(params *quota.ListQuotasParams, authInfo runtime.ClientAuthInfoWriter) (*quota.ListQuotasOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *quota.ListQuotasOK
	if rf, ok := ret.Get(0).(func(*quota.ListQuotasParams, runtime.ClientAuthInfoWriter) *quota.ListQuotasOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*quota.ListQuotasOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*quota.ListQuotasParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockQuotaClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// UpdateQuota provides a mock function with given fields: params, authInfo
func (_m *MockQuotaClientService) UpdateQuota(params *quota.UpdateQuotaParams, authInfo runtime.ClientAuthInfoWriter) (*quota.UpdateQuotaOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *quota.UpdateQuotaOK
	if rf, ok := ret.Get(0).(func(*quota.UpdateQuotaParams, runtime.ClientAuthInfoWriter) *quota.UpdateQuotaOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*quota.UpdateQuotaOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*quota.UpdateQuotaParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockQuotaClientService creates a new instance of MockQuotaClientService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockQuotaClientService(t testing.TB) *MockQuotaClientService {
	mock := &MockQuotaClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
