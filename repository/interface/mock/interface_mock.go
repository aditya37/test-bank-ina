// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package _interface is a generated GoMock package.
package _interface

import (
	context "context"
	reflect "reflect"
	time "time"

	model "github.com/adity37/task/model"
	gomock "github.com/golang/mock/gomock"
	oauth2 "golang.org/x/oauth2"
)

// MockDBReaderWriter is a mock of DBReaderWriter interface.
type MockDBReaderWriter struct {
	ctrl     *gomock.Controller
	recorder *MockDBReaderWriterMockRecorder
}

// MockDBReaderWriterMockRecorder is the mock recorder for MockDBReaderWriter.
type MockDBReaderWriterMockRecorder struct {
	mock *MockDBReaderWriter
}

// NewMockDBReaderWriter creates a new mock instance.
func NewMockDBReaderWriter(ctrl *gomock.Controller) *MockDBReaderWriter {
	mock := &MockDBReaderWriter{ctrl: ctrl}
	mock.recorder = &MockDBReaderWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBReaderWriter) EXPECT() *MockDBReaderWriterMockRecorder {
	return m.recorder
}

// AddTask mocks base method.
func (m *MockDBReaderWriter) AddTask(ctx context.Context, data model.RequestAddTask) (model.ResponseAddTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTask", ctx, data)
	ret0, _ := ret[0].(model.ResponseAddTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTask indicates an expected call of AddTask.
func (mr *MockDBReaderWriterMockRecorder) AddTask(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTask", reflect.TypeOf((*MockDBReaderWriter)(nil).AddTask), ctx, data)
}

// Close mocks base method.
func (m *MockDBReaderWriter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDBReaderWriterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDBReaderWriter)(nil).Close))
}

// DeleteTaskByID mocks base method.
func (m *MockDBReaderWriter) DeleteTaskByID(ctx context.Context, data model.RequestDeleteTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTaskByID", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTaskByID indicates an expected call of DeleteTaskByID.
func (mr *MockDBReaderWriterMockRecorder) DeleteTaskByID(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTaskByID", reflect.TypeOf((*MockDBReaderWriter)(nil).DeleteTaskByID), ctx, data)
}

// DeleteUserByID mocks base method.
func (m *MockDBReaderWriter) DeleteUserByID(ctx context.Context, userid int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserByID", ctx, userid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserByID indicates an expected call of DeleteUserByID.
func (mr *MockDBReaderWriterMockRecorder) DeleteUserByID(ctx, userid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserByID", reflect.TypeOf((*MockDBReaderWriter)(nil).DeleteUserByID), ctx, userid)
}

// FetchTask mocks base method.
func (m *MockDBReaderWriter) FetchTask(ctx context.Context) (model.ResponseFetchTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchTask", ctx)
	ret0, _ := ret[0].(model.ResponseFetchTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchTask indicates an expected call of FetchTask.
func (mr *MockDBReaderWriterMockRecorder) FetchTask(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchTask", reflect.TypeOf((*MockDBReaderWriter)(nil).FetchTask), ctx)
}

// GetTaskByID mocks base method.
func (m *MockDBReaderWriter) GetTaskByID(ctx context.Context, data model.RequestGetTaskByID) (model.ResponseGetTaskById, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskByID", ctx, data)
	ret0, _ := ret[0].(model.ResponseGetTaskById)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskByID indicates an expected call of GetTaskByID.
func (mr *MockDBReaderWriterMockRecorder) GetTaskByID(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskByID", reflect.TypeOf((*MockDBReaderWriter)(nil).GetTaskByID), ctx, data)
}

// GetUserByEmail mocks base method.
func (m *MockDBReaderWriter) GetUserByEmail(ctx context.Context, data model.RequestRegisterUser) (model.ResponseGetUserByID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, data)
	ret0, _ := ret[0].(model.ResponseGetUserByID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockDBReaderWriterMockRecorder) GetUserByEmail(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockDBReaderWriter)(nil).GetUserByEmail), ctx, data)
}

// GetUserByID mocks base method.
func (m *MockDBReaderWriter) GetUserByID(ctx context.Context, data model.RequestGetUserByID) (model.ResponseGetUserByID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, data)
	ret0, _ := ret[0].(model.ResponseGetUserByID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockDBReaderWriterMockRecorder) GetUserByID(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockDBReaderWriter)(nil).GetUserByID), ctx, data)
}

// RegisterUser mocks base method.
func (m *MockDBReaderWriter) RegisterUser(ctx context.Context, data model.RequestRegisterUser) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, data)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockDBReaderWriterMockRecorder) RegisterUser(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockDBReaderWriter)(nil).RegisterUser), ctx, data)
}

// UpdateTask mocks base method.
func (m *MockDBReaderWriter) UpdateTask(ctx context.Context, data model.RequestUpdateTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockDBReaderWriterMockRecorder) UpdateTask(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockDBReaderWriter)(nil).UpdateTask), ctx, data)
}

// UpdateUserByID mocks base method.
func (m *MockDBReaderWriter) UpdateUserByID(ctx context.Context, data model.RequestUpdateUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserByID", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserByID indicates an expected call of UpdateUserByID.
func (mr *MockDBReaderWriterMockRecorder) UpdateUserByID(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserByID", reflect.TypeOf((*MockDBReaderWriter)(nil).UpdateUserByID), ctx, data)
}

// MockRedisReaderWriter is a mock of RedisReaderWriter interface.
type MockRedisReaderWriter struct {
	ctrl     *gomock.Controller
	recorder *MockRedisReaderWriterMockRecorder
}

// MockRedisReaderWriterMockRecorder is the mock recorder for MockRedisReaderWriter.
type MockRedisReaderWriterMockRecorder struct {
	mock *MockRedisReaderWriter
}

// NewMockRedisReaderWriter creates a new mock instance.
func NewMockRedisReaderWriter(ctrl *gomock.Controller) *MockRedisReaderWriter {
	mock := &MockRedisReaderWriter{ctrl: ctrl}
	mock.recorder = &MockRedisReaderWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisReaderWriter) EXPECT() *MockRedisReaderWriterMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRedisReaderWriter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRedisReaderWriterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRedisReaderWriter)(nil).Close))
}

// GetUserSession mocks base method.
func (m *MockRedisReaderWriter) GetUserSession(ctx context.Context, key string) (model.SessionPayload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSession", ctx, key)
	ret0, _ := ret[0].(model.SessionPayload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSession indicates an expected call of GetUserSession.
func (mr *MockRedisReaderWriterMockRecorder) GetUserSession(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSession", reflect.TypeOf((*MockRedisReaderWriter)(nil).GetUserSession), ctx, key)
}

// Set mocks base method.
func (m *MockRedisReaderWriter) Set(ctx context.Context, key string, payload interface{}, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, payload, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockRedisReaderWriterMockRecorder) Set(ctx, key, payload, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedisReaderWriter)(nil).Set), ctx, key, payload, ttl)
}

// MockAuth is a mock of Auth interface.
type MockAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMockRecorder
}

// MockAuthMockRecorder is the mock recorder for MockAuth.
type MockAuthMockRecorder struct {
	mock *MockAuth
}

// NewMockAuth creates a new mock instance.
func NewMockAuth(ctrl *gomock.Controller) *MockAuth {
	mock := &MockAuth{ctrl: ctrl}
	mock.recorder = &MockAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuth) EXPECT() *MockAuthMockRecorder {
	return m.recorder
}

// AuthCodeURL mocks base method.
func (m *MockAuth) AuthCodeURL(state string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthCodeURL", state)
	ret0, _ := ret[0].(string)
	return ret0
}

// AuthCodeURL indicates an expected call of AuthCodeURL.
func (mr *MockAuthMockRecorder) AuthCodeURL(state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthCodeURL", reflect.TypeOf((*MockAuth)(nil).AuthCodeURL), state)
}

// OauthExchange mocks base method.
func (m *MockAuth) OauthExchange(ctx context.Context, code string) (*oauth2.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OauthExchange", ctx, code)
	ret0, _ := ret[0].(*oauth2.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OauthExchange indicates an expected call of OauthExchange.
func (mr *MockAuthMockRecorder) OauthExchange(ctx, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OauthExchange", reflect.TypeOf((*MockAuth)(nil).OauthExchange), ctx, code)
}

// ParseTokenDetail mocks base method.
func (m *MockAuth) ParseTokenDetail(token string) (model.ResponseParseToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseTokenDetail", token)
	ret0, _ := ret[0].(model.ResponseParseToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseTokenDetail indicates an expected call of ParseTokenDetail.
func (mr *MockAuthMockRecorder) ParseTokenDetail(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseTokenDetail", reflect.TypeOf((*MockAuth)(nil).ParseTokenDetail), token)
}