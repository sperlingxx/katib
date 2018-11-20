// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeflow/katib/pkg/db (interfaces: VizierDBInterface)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/kubeflow/katib/pkg/api"
	db "github.com/kubeflow/katib/pkg/db"
	reflect "reflect"
	time "time"
)

// MockVizierDBInterface is a mock of VizierDBInterface interface
type MockVizierDBInterface struct {
	ctrl     *gomock.Controller
	recorder *MockVizierDBInterfaceMockRecorder
}

// MockVizierDBInterfaceMockRecorder is the mock recorder for MockVizierDBInterface
type MockVizierDBInterfaceMockRecorder struct {
	mock *MockVizierDBInterface
}

// NewMockVizierDBInterface creates a new mock instance
func NewMockVizierDBInterface(ctrl *gomock.Controller) *MockVizierDBInterface {
	mock := &MockVizierDBInterface{ctrl: ctrl}
	mock.recorder = &MockVizierDBInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVizierDBInterface) EXPECT() *MockVizierDBInterfaceMockRecorder {
	return m.recorder
}

// CreateStudy mocks base method
func (m *MockVizierDBInterface) CreateStudy(arg0 *api.StudyConfig) (string, error) {
	ret := m.ctrl.Call(m, "CreateStudy", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStudy indicates an expected call of CreateStudy
func (mr *MockVizierDBInterfaceMockRecorder) CreateStudy(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStudy", reflect.TypeOf((*MockVizierDBInterface)(nil).CreateStudy), arg0)
}

// CreateTrial mocks base method
func (m *MockVizierDBInterface) CreateTrial(arg0 *api.Trial) error {
	ret := m.ctrl.Call(m, "CreateTrial", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTrial indicates an expected call of CreateTrial
func (mr *MockVizierDBInterfaceMockRecorder) CreateTrial(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrial", reflect.TypeOf((*MockVizierDBInterface)(nil).CreateTrial), arg0)
}

// CreateWorker mocks base method
func (m *MockVizierDBInterface) CreateWorker(arg0 *api.Worker) (string, error) {
	ret := m.ctrl.Call(m, "CreateWorker", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWorker indicates an expected call of CreateWorker
func (mr *MockVizierDBInterfaceMockRecorder) CreateWorker(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorker", reflect.TypeOf((*MockVizierDBInterface)(nil).CreateWorker), arg0)
}

// DBInit mocks base method
func (m *MockVizierDBInterface) DBInit() {
	m.ctrl.Call(m, "DBInit")
}

// DBInit indicates an expected call of DBInit
func (mr *MockVizierDBInterfaceMockRecorder) DBInit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBInit", reflect.TypeOf((*MockVizierDBInterface)(nil).DBInit))
}

// DeleteStudy mocks base method
func (m *MockVizierDBInterface) DeleteStudy(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteStudy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStudy indicates an expected call of DeleteStudy
func (mr *MockVizierDBInterfaceMockRecorder) DeleteStudy(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStudy", reflect.TypeOf((*MockVizierDBInterface)(nil).DeleteStudy), arg0)
}

// DeleteTrial mocks base method
func (m *MockVizierDBInterface) DeleteTrial(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteTrial", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrial indicates an expected call of DeleteTrial
func (mr *MockVizierDBInterfaceMockRecorder) DeleteTrial(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrial", reflect.TypeOf((*MockVizierDBInterface)(nil).DeleteTrial), arg0)
}

// DeleteWorker mocks base method
func (m *MockVizierDBInterface) DeleteWorker(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteWorker", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWorker indicates an expected call of DeleteWorker
func (mr *MockVizierDBInterfaceMockRecorder) DeleteWorker(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorker", reflect.TypeOf((*MockVizierDBInterface)(nil).DeleteWorker), arg0)
}

// GetEarlyStopParam mocks base method
func (m *MockVizierDBInterface) GetEarlyStopParam(arg0 string) ([]*api.EarlyStoppingParameter, error) {
	ret := m.ctrl.Call(m, "GetEarlyStopParam", arg0)
	ret0, _ := ret[0].([]*api.EarlyStoppingParameter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEarlyStopParam indicates an expected call of GetEarlyStopParam
func (mr *MockVizierDBInterfaceMockRecorder) GetEarlyStopParam(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEarlyStopParam", reflect.TypeOf((*MockVizierDBInterface)(nil).GetEarlyStopParam), arg0)
}

// GetEarlyStopParamList mocks base method
func (m *MockVizierDBInterface) GetEarlyStopParamList(arg0 string) ([]*api.EarlyStoppingParameterSet, error) {
	ret := m.ctrl.Call(m, "GetEarlyStopParamList", arg0)
	ret0, _ := ret[0].([]*api.EarlyStoppingParameterSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEarlyStopParamList indicates an expected call of GetEarlyStopParamList
func (mr *MockVizierDBInterfaceMockRecorder) GetEarlyStopParamList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEarlyStopParamList", reflect.TypeOf((*MockVizierDBInterface)(nil).GetEarlyStopParamList), arg0)
}

// GetStudyConfig mocks base method
func (m *MockVizierDBInterface) GetStudyConfig(arg0 string) (*api.StudyConfig, error) {
	ret := m.ctrl.Call(m, "GetStudyConfig", arg0)
	ret0, _ := ret[0].(*api.StudyConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudyConfig indicates an expected call of GetStudyConfig
func (mr *MockVizierDBInterfaceMockRecorder) GetStudyConfig(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudyConfig", reflect.TypeOf((*MockVizierDBInterface)(nil).GetStudyConfig), arg0)
}

// GetStudyList mocks base method
func (m *MockVizierDBInterface) GetStudyList() ([]string, error) {
	ret := m.ctrl.Call(m, "GetStudyList")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudyList indicates an expected call of GetStudyList
func (mr *MockVizierDBInterfaceMockRecorder) GetStudyList() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudyList", reflect.TypeOf((*MockVizierDBInterface)(nil).GetStudyList))
}

// GetSuggestionParam mocks base method
func (m *MockVizierDBInterface) GetSuggestionParam(arg0 string) ([]*api.SuggestionParameter, error) {
	ret := m.ctrl.Call(m, "GetSuggestionParam", arg0)
	ret0, _ := ret[0].([]*api.SuggestionParameter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSuggestionParam indicates an expected call of GetSuggestionParam
func (mr *MockVizierDBInterfaceMockRecorder) GetSuggestionParam(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSuggestionParam", reflect.TypeOf((*MockVizierDBInterface)(nil).GetSuggestionParam), arg0)
}

// GetSuggestionParamList mocks base method
func (m *MockVizierDBInterface) GetSuggestionParamList(arg0 string) ([]*api.SuggestionParameterSet, error) {
	ret := m.ctrl.Call(m, "GetSuggestionParamList", arg0)
	ret0, _ := ret[0].([]*api.SuggestionParameterSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSuggestionParamList indicates an expected call of GetSuggestionParamList
func (mr *MockVizierDBInterfaceMockRecorder) GetSuggestionParamList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSuggestionParamList", reflect.TypeOf((*MockVizierDBInterface)(nil).GetSuggestionParamList), arg0)
}

// GetTrial mocks base method
func (m *MockVizierDBInterface) GetTrial(arg0 string) (*api.Trial, error) {
	ret := m.ctrl.Call(m, "GetTrial", arg0)
	ret0, _ := ret[0].(*api.Trial)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrial indicates an expected call of GetTrial
func (mr *MockVizierDBInterfaceMockRecorder) GetTrial(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrial", reflect.TypeOf((*MockVizierDBInterface)(nil).GetTrial), arg0)
}

// GetTrialList mocks base method
func (m *MockVizierDBInterface) GetTrialList(arg0 string) ([]*api.Trial, error) {
	ret := m.ctrl.Call(m, "GetTrialList", arg0)
	ret0, _ := ret[0].([]*api.Trial)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrialList indicates an expected call of GetTrialList
func (mr *MockVizierDBInterfaceMockRecorder) GetTrialList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrialList", reflect.TypeOf((*MockVizierDBInterface)(nil).GetTrialList), arg0)
}

// GetWorker mocks base method
func (m *MockVizierDBInterface) GetWorker(arg0 string) (*api.Worker, error) {
	ret := m.ctrl.Call(m, "GetWorker", arg0)
	ret0, _ := ret[0].(*api.Worker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorker indicates an expected call of GetWorker
func (mr *MockVizierDBInterfaceMockRecorder) GetWorker(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorker", reflect.TypeOf((*MockVizierDBInterface)(nil).GetWorker), arg0)
}

// GetWorkerFullInfo mocks base method
func (m *MockVizierDBInterface) GetWorkerFullInfo(arg0, arg1, arg2 string, arg3 bool) (*api.GetWorkerFullInfoReply, error) {
	ret := m.ctrl.Call(m, "GetWorkerFullInfo", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*api.GetWorkerFullInfoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkerFullInfo indicates an expected call of GetWorkerFullInfo
func (mr *MockVizierDBInterfaceMockRecorder) GetWorkerFullInfo(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkerFullInfo", reflect.TypeOf((*MockVizierDBInterface)(nil).GetWorkerFullInfo), arg0, arg1, arg2, arg3)
}

// GetWorkerList mocks base method
func (m *MockVizierDBInterface) GetWorkerList(arg0, arg1 string) ([]*api.Worker, error) {
	ret := m.ctrl.Call(m, "GetWorkerList", arg0, arg1)
	ret0, _ := ret[0].([]*api.Worker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkerList indicates an expected call of GetWorkerList
func (mr *MockVizierDBInterfaceMockRecorder) GetWorkerList(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkerList", reflect.TypeOf((*MockVizierDBInterface)(nil).GetWorkerList), arg0, arg1)
}

// GetWorkerLogs mocks base method
func (m *MockVizierDBInterface) GetWorkerLogs(arg0 string, arg1 *db.GetWorkerLogOpts) ([]*db.WorkerLog, error) {
	ret := m.ctrl.Call(m, "GetWorkerLogs", arg0, arg1)
	ret0, _ := ret[0].([]*db.WorkerLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkerLogs indicates an expected call of GetWorkerLogs
func (mr *MockVizierDBInterfaceMockRecorder) GetWorkerLogs(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkerLogs", reflect.TypeOf((*MockVizierDBInterface)(nil).GetWorkerLogs), arg0, arg1)
}

// GetWorkerStatus mocks base method
func (m *MockVizierDBInterface) GetWorkerStatus(arg0 string) (*api.State, error) {
	ret := m.ctrl.Call(m, "GetWorkerStatus", arg0)
	ret0, _ := ret[0].(*api.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkerStatus indicates an expected call of GetWorkerStatus
func (mr *MockVizierDBInterfaceMockRecorder) GetWorkerStatus(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkerStatus", reflect.TypeOf((*MockVizierDBInterface)(nil).GetWorkerStatus), arg0)
}

// GetWorkerTimestamp mocks base method
func (m *MockVizierDBInterface) GetWorkerTimestamp(arg0 string) (*time.Time, error) {
	ret := m.ctrl.Call(m, "GetWorkerTimestamp", arg0)
	ret0, _ := ret[0].(*time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkerTimestamp indicates an expected call of GetWorkerTimestamp
func (mr *MockVizierDBInterfaceMockRecorder) GetWorkerTimestamp(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkerTimestamp", reflect.TypeOf((*MockVizierDBInterface)(nil).GetWorkerTimestamp), arg0)
}

// SetEarlyStopParam mocks base method
func (m *MockVizierDBInterface) SetEarlyStopParam(arg0, arg1 string, arg2 []*api.EarlyStoppingParameter) (string, error) {
	ret := m.ctrl.Call(m, "SetEarlyStopParam", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetEarlyStopParam indicates an expected call of SetEarlyStopParam
func (mr *MockVizierDBInterfaceMockRecorder) SetEarlyStopParam(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEarlyStopParam", reflect.TypeOf((*MockVizierDBInterface)(nil).SetEarlyStopParam), arg0, arg1, arg2)
}

// SetSuggestionParam mocks base method
func (m *MockVizierDBInterface) SetSuggestionParam(arg0, arg1 string, arg2 []*api.SuggestionParameter) (string, error) {
	ret := m.ctrl.Call(m, "SetSuggestionParam", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetSuggestionParam indicates an expected call of SetSuggestionParam
func (mr *MockVizierDBInterfaceMockRecorder) SetSuggestionParam(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSuggestionParam", reflect.TypeOf((*MockVizierDBInterface)(nil).SetSuggestionParam), arg0, arg1, arg2)
}

// StoreWorkerLogs mocks base method
func (m *MockVizierDBInterface) StoreWorkerLogs(arg0 string, arg1 []*api.MetricsLog) error {
	ret := m.ctrl.Call(m, "StoreWorkerLogs", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreWorkerLogs indicates an expected call of StoreWorkerLogs
func (mr *MockVizierDBInterfaceMockRecorder) StoreWorkerLogs(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreWorkerLogs", reflect.TypeOf((*MockVizierDBInterface)(nil).StoreWorkerLogs), arg0, arg1)
}

// UpdateEarlyStopParam mocks base method
func (m *MockVizierDBInterface) UpdateEarlyStopParam(arg0 string, arg1 []*api.EarlyStoppingParameter) error {
	ret := m.ctrl.Call(m, "UpdateEarlyStopParam", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEarlyStopParam indicates an expected call of UpdateEarlyStopParam
func (mr *MockVizierDBInterfaceMockRecorder) UpdateEarlyStopParam(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEarlyStopParam", reflect.TypeOf((*MockVizierDBInterface)(nil).UpdateEarlyStopParam), arg0, arg1)
}

// UpdateSuggestionParam mocks base method
func (m *MockVizierDBInterface) UpdateSuggestionParam(arg0 string, arg1 []*api.SuggestionParameter) error {
	ret := m.ctrl.Call(m, "UpdateSuggestionParam", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSuggestionParam indicates an expected call of UpdateSuggestionParam
func (mr *MockVizierDBInterfaceMockRecorder) UpdateSuggestionParam(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSuggestionParam", reflect.TypeOf((*MockVizierDBInterface)(nil).UpdateSuggestionParam), arg0, arg1)
}

// UpdateWorker mocks base method
func (m *MockVizierDBInterface) UpdateWorker(arg0 string, arg1 api.State) error {
	ret := m.ctrl.Call(m, "UpdateWorker", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWorker indicates an expected call of UpdateWorker
func (mr *MockVizierDBInterfaceMockRecorder) UpdateWorker(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorker", reflect.TypeOf((*MockVizierDBInterface)(nil).UpdateWorker), arg0, arg1)
}