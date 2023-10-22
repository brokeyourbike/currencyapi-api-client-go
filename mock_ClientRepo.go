// Code generated by mockery v2.33.0. DO NOT EDIT.

package currencyapi

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockClientRepo is an autogenerated mock type for the ClientRepo type
type MockClientRepo struct {
	mock.Mock
}

// GetHistoricalRate provides a mock function with given fields: ctx, baseCurrency, currencies, date
func (_m *MockClientRepo) GetHistoricalRate(ctx context.Context, baseCurrency string, currencies []string, date time.Time) (RateResponse, error) {
	ret := _m.Called(ctx, baseCurrency, currencies, date)

	var r0 RateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, time.Time) (RateResponse, error)); ok {
		return rf(ctx, baseCurrency, currencies, date)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, time.Time) RateResponse); ok {
		r0 = rf(ctx, baseCurrency, currencies, date)
	} else {
		r0 = ret.Get(0).(RateResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []string, time.Time) error); ok {
		r1 = rf(ctx, baseCurrency, currencies, date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestRate provides a mock function with given fields: ctx, baseCurrency, currencies
func (_m *MockClientRepo) GetLatestRate(ctx context.Context, baseCurrency string, currencies []string) (RateResponse, error) {
	ret := _m.Called(ctx, baseCurrency, currencies)

	var r0 RateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) (RateResponse, error)); ok {
		return rf(ctx, baseCurrency, currencies)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) RateResponse); ok {
		r0 = rf(ctx, baseCurrency, currencies)
	} else {
		r0 = ret.Get(0).(RateResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []string) error); ok {
		r1 = rf(ctx, baseCurrency, currencies)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockClientRepo creates a new instance of MockClientRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClientRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClientRepo {
	mock := &MockClientRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
