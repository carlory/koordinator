/*
Copyright 2022 The Koordinator Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: apis/runtime/v1alpha1/api_grpc.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/koordinator-sh/koordinator/apis/runtime/v1alpha1"
	grpc "google.golang.org/grpc"
)

// MockRuntimeHookServiceClient is a mock of RuntimeHookServiceClient interface.
type MockRuntimeHookServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockRuntimeHookServiceClientMockRecorder
}

// MockRuntimeHookServiceClientMockRecorder is the mock recorder for MockRuntimeHookServiceClient.
type MockRuntimeHookServiceClientMockRecorder struct {
	mock *MockRuntimeHookServiceClient
}

// NewMockRuntimeHookServiceClient creates a new mock instance.
func NewMockRuntimeHookServiceClient(ctrl *gomock.Controller) *MockRuntimeHookServiceClient {
	mock := &MockRuntimeHookServiceClient{ctrl: ctrl}
	mock.recorder = &MockRuntimeHookServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRuntimeHookServiceClient) EXPECT() *MockRuntimeHookServiceClientMockRecorder {
	return m.recorder
}

// PostStartContainerHook mocks base method.
func (m *MockRuntimeHookServiceClient) PostStartContainerHook(ctx context.Context, in *v1alpha1.ContainerResourceHookRequest, opts ...grpc.CallOption) (*v1alpha1.ContainerResourceHookResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostStartContainerHook", varargs...)
	ret0, _ := ret[0].(*v1alpha1.ContainerResourceHookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostStartContainerHook indicates an expected call of PostStartContainerHook.
func (mr *MockRuntimeHookServiceClientMockRecorder) PostStartContainerHook(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostStartContainerHook", reflect.TypeOf((*MockRuntimeHookServiceClient)(nil).PostStartContainerHook), varargs...)
}

// PostStopContainerHook mocks base method.
func (m *MockRuntimeHookServiceClient) PostStopContainerHook(ctx context.Context, in *v1alpha1.ContainerResourceHookRequest, opts ...grpc.CallOption) (*v1alpha1.ContainerResourceHookResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostStopContainerHook", varargs...)
	ret0, _ := ret[0].(*v1alpha1.ContainerResourceHookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostStopContainerHook indicates an expected call of PostStopContainerHook.
func (mr *MockRuntimeHookServiceClientMockRecorder) PostStopContainerHook(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostStopContainerHook", reflect.TypeOf((*MockRuntimeHookServiceClient)(nil).PostStopContainerHook), varargs...)
}

// PostStopPodSandboxHook mocks base method.
func (m *MockRuntimeHookServiceClient) PostStopPodSandboxHook(ctx context.Context, in *v1alpha1.PodSandboxHookRequest, opts ...grpc.CallOption) (*v1alpha1.PodSandboxHookResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostStopPodSandboxHook", varargs...)
	ret0, _ := ret[0].(*v1alpha1.PodSandboxHookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostStopPodSandboxHook indicates an expected call of PostStopPodSandboxHook.
func (mr *MockRuntimeHookServiceClientMockRecorder) PostStopPodSandboxHook(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostStopPodSandboxHook", reflect.TypeOf((*MockRuntimeHookServiceClient)(nil).PostStopPodSandboxHook), varargs...)
}

// PreCreateContainerHook mocks base method.
func (m *MockRuntimeHookServiceClient) PreCreateContainerHook(ctx context.Context, in *v1alpha1.ContainerResourceHookRequest, opts ...grpc.CallOption) (*v1alpha1.ContainerResourceHookResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PreCreateContainerHook", varargs...)
	ret0, _ := ret[0].(*v1alpha1.ContainerResourceHookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PreCreateContainerHook indicates an expected call of PreCreateContainerHook.
func (mr *MockRuntimeHookServiceClientMockRecorder) PreCreateContainerHook(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreCreateContainerHook", reflect.TypeOf((*MockRuntimeHookServiceClient)(nil).PreCreateContainerHook), varargs...)
}

// PreRunPodSandboxHook mocks base method.
func (m *MockRuntimeHookServiceClient) PreRunPodSandboxHook(ctx context.Context, in *v1alpha1.PodSandboxHookRequest, opts ...grpc.CallOption) (*v1alpha1.PodSandboxHookResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PreRunPodSandboxHook", varargs...)
	ret0, _ := ret[0].(*v1alpha1.PodSandboxHookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PreRunPodSandboxHook indicates an expected call of PreRunPodSandboxHook.
func (mr *MockRuntimeHookServiceClientMockRecorder) PreRunPodSandboxHook(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreRunPodSandboxHook", reflect.TypeOf((*MockRuntimeHookServiceClient)(nil).PreRunPodSandboxHook), varargs...)
}

// PreStartContainerHook mocks base method.
func (m *MockRuntimeHookServiceClient) PreStartContainerHook(ctx context.Context, in *v1alpha1.ContainerResourceHookRequest, opts ...grpc.CallOption) (*v1alpha1.ContainerResourceHookResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PreStartContainerHook", varargs...)
	ret0, _ := ret[0].(*v1alpha1.ContainerResourceHookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PreStartContainerHook indicates an expected call of PreStartContainerHook.
func (mr *MockRuntimeHookServiceClientMockRecorder) PreStartContainerHook(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreStartContainerHook", reflect.TypeOf((*MockRuntimeHookServiceClient)(nil).PreStartContainerHook), varargs...)
}

// PreUpdateContainerResourcesHook mocks base method.
func (m *MockRuntimeHookServiceClient) PreUpdateContainerResourcesHook(ctx context.Context, in *v1alpha1.ContainerResourceHookRequest, opts ...grpc.CallOption) (*v1alpha1.ContainerResourceHookResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PreUpdateContainerResourcesHook", varargs...)
	ret0, _ := ret[0].(*v1alpha1.ContainerResourceHookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PreUpdateContainerResourcesHook indicates an expected call of PreUpdateContainerResourcesHook.
func (mr *MockRuntimeHookServiceClientMockRecorder) PreUpdateContainerResourcesHook(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreUpdateContainerResourcesHook", reflect.TypeOf((*MockRuntimeHookServiceClient)(nil).PreUpdateContainerResourcesHook), varargs...)
}

// MockRuntimeHookServiceServer is a mock of RuntimeHookServiceServer interface.
type MockRuntimeHookServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockRuntimeHookServiceServerMockRecorder
}

// MockRuntimeHookServiceServerMockRecorder is the mock recorder for MockRuntimeHookServiceServer.
type MockRuntimeHookServiceServerMockRecorder struct {
	mock *MockRuntimeHookServiceServer
}

// NewMockRuntimeHookServiceServer creates a new mock instance.
func NewMockRuntimeHookServiceServer(ctrl *gomock.Controller) *MockRuntimeHookServiceServer {
	mock := &MockRuntimeHookServiceServer{ctrl: ctrl}
	mock.recorder = &MockRuntimeHookServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRuntimeHookServiceServer) EXPECT() *MockRuntimeHookServiceServerMockRecorder {
	return m.recorder
}

// PostStartContainerHook mocks base method.
func (m *MockRuntimeHookServiceServer) PostStartContainerHook(arg0 context.Context, arg1 *v1alpha1.ContainerResourceHookRequest) (*v1alpha1.ContainerResourceHookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostStartContainerHook", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.ContainerResourceHookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostStartContainerHook indicates an expected call of PostStartContainerHook.
func (mr *MockRuntimeHookServiceServerMockRecorder) PostStartContainerHook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostStartContainerHook", reflect.TypeOf((*MockRuntimeHookServiceServer)(nil).PostStartContainerHook), arg0, arg1)
}

// PostStopContainerHook mocks base method.
func (m *MockRuntimeHookServiceServer) PostStopContainerHook(arg0 context.Context, arg1 *v1alpha1.ContainerResourceHookRequest) (*v1alpha1.ContainerResourceHookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostStopContainerHook", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.ContainerResourceHookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostStopContainerHook indicates an expected call of PostStopContainerHook.
func (mr *MockRuntimeHookServiceServerMockRecorder) PostStopContainerHook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostStopContainerHook", reflect.TypeOf((*MockRuntimeHookServiceServer)(nil).PostStopContainerHook), arg0, arg1)
}

// PostStopPodSandboxHook mocks base method.
func (m *MockRuntimeHookServiceServer) PostStopPodSandboxHook(arg0 context.Context, arg1 *v1alpha1.PodSandboxHookRequest) (*v1alpha1.PodSandboxHookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostStopPodSandboxHook", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.PodSandboxHookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostStopPodSandboxHook indicates an expected call of PostStopPodSandboxHook.
func (mr *MockRuntimeHookServiceServerMockRecorder) PostStopPodSandboxHook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostStopPodSandboxHook", reflect.TypeOf((*MockRuntimeHookServiceServer)(nil).PostStopPodSandboxHook), arg0, arg1)
}

// PreCreateContainerHook mocks base method.
func (m *MockRuntimeHookServiceServer) PreCreateContainerHook(arg0 context.Context, arg1 *v1alpha1.ContainerResourceHookRequest) (*v1alpha1.ContainerResourceHookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreCreateContainerHook", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.ContainerResourceHookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PreCreateContainerHook indicates an expected call of PreCreateContainerHook.
func (mr *MockRuntimeHookServiceServerMockRecorder) PreCreateContainerHook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreCreateContainerHook", reflect.TypeOf((*MockRuntimeHookServiceServer)(nil).PreCreateContainerHook), arg0, arg1)
}

// PreRunPodSandboxHook mocks base method.
func (m *MockRuntimeHookServiceServer) PreRunPodSandboxHook(arg0 context.Context, arg1 *v1alpha1.PodSandboxHookRequest) (*v1alpha1.PodSandboxHookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreRunPodSandboxHook", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.PodSandboxHookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PreRunPodSandboxHook indicates an expected call of PreRunPodSandboxHook.
func (mr *MockRuntimeHookServiceServerMockRecorder) PreRunPodSandboxHook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreRunPodSandboxHook", reflect.TypeOf((*MockRuntimeHookServiceServer)(nil).PreRunPodSandboxHook), arg0, arg1)
}

// PreStartContainerHook mocks base method.
func (m *MockRuntimeHookServiceServer) PreStartContainerHook(arg0 context.Context, arg1 *v1alpha1.ContainerResourceHookRequest) (*v1alpha1.ContainerResourceHookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreStartContainerHook", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.ContainerResourceHookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PreStartContainerHook indicates an expected call of PreStartContainerHook.
func (mr *MockRuntimeHookServiceServerMockRecorder) PreStartContainerHook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreStartContainerHook", reflect.TypeOf((*MockRuntimeHookServiceServer)(nil).PreStartContainerHook), arg0, arg1)
}

// PreUpdateContainerResourcesHook mocks base method.
func (m *MockRuntimeHookServiceServer) PreUpdateContainerResourcesHook(arg0 context.Context, arg1 *v1alpha1.ContainerResourceHookRequest) (*v1alpha1.ContainerResourceHookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreUpdateContainerResourcesHook", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.ContainerResourceHookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PreUpdateContainerResourcesHook indicates an expected call of PreUpdateContainerResourcesHook.
func (mr *MockRuntimeHookServiceServerMockRecorder) PreUpdateContainerResourcesHook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreUpdateContainerResourcesHook", reflect.TypeOf((*MockRuntimeHookServiceServer)(nil).PreUpdateContainerResourcesHook), arg0, arg1)
}

// mustEmbedUnimplementedRuntimeHookServiceServer mocks base method.
func (m *MockRuntimeHookServiceServer) mustEmbedUnimplementedRuntimeHookServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedRuntimeHookServiceServer")
}

// mustEmbedUnimplementedRuntimeHookServiceServer indicates an expected call of mustEmbedUnimplementedRuntimeHookServiceServer.
func (mr *MockRuntimeHookServiceServerMockRecorder) mustEmbedUnimplementedRuntimeHookServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedRuntimeHookServiceServer", reflect.TypeOf((*MockRuntimeHookServiceServer)(nil).mustEmbedUnimplementedRuntimeHookServiceServer))
}

// MockUnsafeRuntimeHookServiceServer is a mock of UnsafeRuntimeHookServiceServer interface.
type MockUnsafeRuntimeHookServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeRuntimeHookServiceServerMockRecorder
}

// MockUnsafeRuntimeHookServiceServerMockRecorder is the mock recorder for MockUnsafeRuntimeHookServiceServer.
type MockUnsafeRuntimeHookServiceServerMockRecorder struct {
	mock *MockUnsafeRuntimeHookServiceServer
}

// NewMockUnsafeRuntimeHookServiceServer creates a new mock instance.
func NewMockUnsafeRuntimeHookServiceServer(ctrl *gomock.Controller) *MockUnsafeRuntimeHookServiceServer {
	mock := &MockUnsafeRuntimeHookServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeRuntimeHookServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeRuntimeHookServiceServer) EXPECT() *MockUnsafeRuntimeHookServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedRuntimeHookServiceServer mocks base method.
func (m *MockUnsafeRuntimeHookServiceServer) mustEmbedUnimplementedRuntimeHookServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedRuntimeHookServiceServer")
}

// mustEmbedUnimplementedRuntimeHookServiceServer indicates an expected call of mustEmbedUnimplementedRuntimeHookServiceServer.
func (mr *MockUnsafeRuntimeHookServiceServerMockRecorder) mustEmbedUnimplementedRuntimeHookServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedRuntimeHookServiceServer", reflect.TypeOf((*MockUnsafeRuntimeHookServiceServer)(nil).mustEmbedUnimplementedRuntimeHookServiceServer))
}
