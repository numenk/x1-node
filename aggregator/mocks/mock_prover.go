// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	pb "github.com/0xPolygonHermez/zkevm-node/aggregator/pb"
	mock "github.com/stretchr/testify/mock"
)

// ProverMock is an autogenerated mock type for the proverInterface type
type ProverMock struct {
	mock.Mock
}

// FinalProof provides a mock function with given fields: inputProof
func (_m *ProverMock) FinalProof(inputProof string) (string, error) {
	ret := _m.Called(inputProof)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(inputProof)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(inputProof)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ID provides a mock function with given fields:
func (_m *ProverMock) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsIdle provides a mock function with given fields:
func (_m *ProverMock) IsIdle() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ProveBatch provides a mock function with given fields: input
func (_m *ProverMock) ProveBatch(input *pb.InputProver) (string, error) {
	ret := _m.Called(input)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pb.InputProver) string); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pb.InputProver) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitFinalProof provides a mock function with given fields: ctx, proofID
func (_m *ProverMock) WaitFinalProof(ctx context.Context, proofID string) (*pb.FinalProof, error) {
	ret := _m.Called(ctx, proofID)

	var r0 *pb.FinalProof
	if rf, ok := ret.Get(0).(func(context.Context, string) *pb.FinalProof); ok {
		r0 = rf(ctx, proofID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.FinalProof)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, proofID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitRecursiveProof provides a mock function with given fields: ctx, proofID
func (_m *ProverMock) WaitRecursiveProof(ctx context.Context, proofID string) (string, error) {
	ret := _m.Called(ctx, proofID)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, proofID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, proofID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProverMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewProverMock creates a new instance of ProverMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProverMock(t mockConstructorTestingTNewProverMock) *ProverMock {
	mock := &ProverMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
