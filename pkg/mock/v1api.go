// Code generated by MockGen. DO NOT EDIT.
// Source: api/prometheus/v1/api.go
//
// Generated by this command:
//
//	mockgen -source api/prometheus/v1/api.go -destination mock.go -package mock API
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	http "net/http"
	url "net/url"
	reflect "reflect"
	time "time"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	model "github.com/prometheus/common/model"
	gomock "go.uber.org/mock/gomock"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// AlertManagers mocks base method.
func (m *MockAPI) AlertManagers(ctx context.Context) (v1.AlertManagersResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertManagers", ctx)
	ret0, _ := ret[0].(v1.AlertManagersResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertManagers indicates an expected call of AlertManagers.
func (mr *MockAPIMockRecorder) AlertManagers(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertManagers", reflect.TypeOf((*MockAPI)(nil).AlertManagers), ctx)
}

// Alerts mocks base method.
func (m *MockAPI) Alerts(ctx context.Context) (v1.AlertsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Alerts", ctx)
	ret0, _ := ret[0].(v1.AlertsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Alerts indicates an expected call of Alerts.
func (mr *MockAPIMockRecorder) Alerts(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Alerts", reflect.TypeOf((*MockAPI)(nil).Alerts), ctx)
}

// Buildinfo mocks base method.
func (m *MockAPI) Buildinfo(ctx context.Context) (v1.BuildinfoResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Buildinfo", ctx)
	ret0, _ := ret[0].(v1.BuildinfoResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Buildinfo indicates an expected call of Buildinfo.
func (mr *MockAPIMockRecorder) Buildinfo(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Buildinfo", reflect.TypeOf((*MockAPI)(nil).Buildinfo), ctx)
}

// CleanTombstones mocks base method.
func (m *MockAPI) CleanTombstones(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanTombstones", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanTombstones indicates an expected call of CleanTombstones.
func (mr *MockAPIMockRecorder) CleanTombstones(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanTombstones", reflect.TypeOf((*MockAPI)(nil).CleanTombstones), ctx)
}

// Config mocks base method.
func (m *MockAPI) Config(ctx context.Context) (v1.ConfigResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config", ctx)
	ret0, _ := ret[0].(v1.ConfigResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Config indicates an expected call of Config.
func (mr *MockAPIMockRecorder) Config(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockAPI)(nil).Config), ctx)
}

// DeleteSeries mocks base method.
func (m *MockAPI) DeleteSeries(ctx context.Context, matches []string, startTime, endTime time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSeries", ctx, matches, startTime, endTime)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSeries indicates an expected call of DeleteSeries.
func (mr *MockAPIMockRecorder) DeleteSeries(ctx, matches, startTime, endTime any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSeries", reflect.TypeOf((*MockAPI)(nil).DeleteSeries), ctx, matches, startTime, endTime)
}

// Flags mocks base method.
func (m *MockAPI) Flags(ctx context.Context) (v1.FlagsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flags", ctx)
	ret0, _ := ret[0].(v1.FlagsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Flags indicates an expected call of Flags.
func (mr *MockAPIMockRecorder) Flags(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flags", reflect.TypeOf((*MockAPI)(nil).Flags), ctx)
}

// LabelNames mocks base method.
func (m *MockAPI) LabelNames(ctx context.Context, matches []string, startTime, endTime time.Time) ([]string, v1.Warnings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LabelNames", ctx, matches, startTime, endTime)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(v1.Warnings)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LabelNames indicates an expected call of LabelNames.
func (mr *MockAPIMockRecorder) LabelNames(ctx, matches, startTime, endTime any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LabelNames", reflect.TypeOf((*MockAPI)(nil).LabelNames), ctx, matches, startTime, endTime)
}

// LabelValues mocks base method.
func (m *MockAPI) LabelValues(ctx context.Context, label string, matches []string, startTime, endTime time.Time) (model.LabelValues, v1.Warnings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LabelValues", ctx, label, matches, startTime, endTime)
	ret0, _ := ret[0].(model.LabelValues)
	ret1, _ := ret[1].(v1.Warnings)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LabelValues indicates an expected call of LabelValues.
func (mr *MockAPIMockRecorder) LabelValues(ctx, label, matches, startTime, endTime any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LabelValues", reflect.TypeOf((*MockAPI)(nil).LabelValues), ctx, label, matches, startTime, endTime)
}

// Metadata mocks base method.
func (m *MockAPI) Metadata(ctx context.Context, metric, limit string) (map[string][]v1.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metadata", ctx, metric, limit)
	ret0, _ := ret[0].(map[string][]v1.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Metadata indicates an expected call of Metadata.
func (mr *MockAPIMockRecorder) Metadata(ctx, metric, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metadata", reflect.TypeOf((*MockAPI)(nil).Metadata), ctx, metric, limit)
}

// Query mocks base method.
func (m *MockAPI) Query(ctx context.Context, query string, ts time.Time, opts ...v1.Option) (model.Value, v1.Warnings, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, query, ts}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(model.Value)
	ret1, _ := ret[1].(v1.Warnings)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Query indicates an expected call of Query.
func (mr *MockAPIMockRecorder) Query(ctx, query, ts any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, query, ts}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockAPI)(nil).Query), varargs...)
}

// QueryExemplars mocks base method.
func (m *MockAPI) QueryExemplars(ctx context.Context, query string, startTime, endTime time.Time) ([]v1.ExemplarQueryResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryExemplars", ctx, query, startTime, endTime)
	ret0, _ := ret[0].([]v1.ExemplarQueryResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryExemplars indicates an expected call of QueryExemplars.
func (mr *MockAPIMockRecorder) QueryExemplars(ctx, query, startTime, endTime any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryExemplars", reflect.TypeOf((*MockAPI)(nil).QueryExemplars), ctx, query, startTime, endTime)
}

// QueryRange mocks base method.
func (m *MockAPI) QueryRange(ctx context.Context, query string, r v1.Range, opts ...v1.Option) (model.Value, v1.Warnings, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, query, r}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRange", varargs...)
	ret0, _ := ret[0].(model.Value)
	ret1, _ := ret[1].(v1.Warnings)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QueryRange indicates an expected call of QueryRange.
func (mr *MockAPIMockRecorder) QueryRange(ctx, query, r any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, query, r}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRange", reflect.TypeOf((*MockAPI)(nil).QueryRange), varargs...)
}

// Rules mocks base method.
func (m *MockAPI) Rules(ctx context.Context) (v1.RulesResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rules", ctx)
	ret0, _ := ret[0].(v1.RulesResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Rules indicates an expected call of Rules.
func (mr *MockAPIMockRecorder) Rules(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rules", reflect.TypeOf((*MockAPI)(nil).Rules), ctx)
}

// Runtimeinfo mocks base method.
func (m *MockAPI) Runtimeinfo(ctx context.Context) (v1.RuntimeinfoResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Runtimeinfo", ctx)
	ret0, _ := ret[0].(v1.RuntimeinfoResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Runtimeinfo indicates an expected call of Runtimeinfo.
func (mr *MockAPIMockRecorder) Runtimeinfo(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Runtimeinfo", reflect.TypeOf((*MockAPI)(nil).Runtimeinfo), ctx)
}

// Series mocks base method.
func (m *MockAPI) Series(ctx context.Context, matches []string, startTime, endTime time.Time) ([]model.LabelSet, v1.Warnings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Series", ctx, matches, startTime, endTime)
	ret0, _ := ret[0].([]model.LabelSet)
	ret1, _ := ret[1].(v1.Warnings)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Series indicates an expected call of Series.
func (mr *MockAPIMockRecorder) Series(ctx, matches, startTime, endTime any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Series", reflect.TypeOf((*MockAPI)(nil).Series), ctx, matches, startTime, endTime)
}

// Snapshot mocks base method.
func (m *MockAPI) Snapshot(ctx context.Context, skipHead bool) (v1.SnapshotResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot", ctx, skipHead)
	ret0, _ := ret[0].(v1.SnapshotResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Snapshot indicates an expected call of Snapshot.
func (mr *MockAPIMockRecorder) Snapshot(ctx, skipHead any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockAPI)(nil).Snapshot), ctx, skipHead)
}

// TSDB mocks base method.
func (m *MockAPI) TSDB(ctx context.Context) (v1.TSDBResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TSDB", ctx)
	ret0, _ := ret[0].(v1.TSDBResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TSDB indicates an expected call of TSDB.
func (mr *MockAPIMockRecorder) TSDB(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TSDB", reflect.TypeOf((*MockAPI)(nil).TSDB), ctx)
}

// Targets mocks base method.
func (m *MockAPI) Targets(ctx context.Context) (v1.TargetsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Targets", ctx)
	ret0, _ := ret[0].(v1.TargetsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Targets indicates an expected call of Targets.
func (mr *MockAPIMockRecorder) Targets(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Targets", reflect.TypeOf((*MockAPI)(nil).Targets), ctx)
}

// TargetsMetadata mocks base method.
func (m *MockAPI) TargetsMetadata(ctx context.Context, matchTarget, metric, limit string) ([]v1.MetricMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TargetsMetadata", ctx, matchTarget, metric, limit)
	ret0, _ := ret[0].([]v1.MetricMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TargetsMetadata indicates an expected call of TargetsMetadata.
func (mr *MockAPIMockRecorder) TargetsMetadata(ctx, matchTarget, metric, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TargetsMetadata", reflect.TypeOf((*MockAPI)(nil).TargetsMetadata), ctx, matchTarget, metric, limit)
}

// WalReplay mocks base method.
func (m *MockAPI) WalReplay(ctx context.Context) (v1.WalReplayStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalReplay", ctx)
	ret0, _ := ret[0].(v1.WalReplayStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalReplay indicates an expected call of WalReplay.
func (mr *MockAPIMockRecorder) WalReplay(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalReplay", reflect.TypeOf((*MockAPI)(nil).WalReplay), ctx)
}

// MockapiClient is a mock of apiClient interface.
type MockapiClient struct {
	ctrl     *gomock.Controller
	recorder *MockapiClientMockRecorder
}

// MockapiClientMockRecorder is the mock recorder for MockapiClient.
type MockapiClientMockRecorder struct {
	mock *MockapiClient
}

// NewMockapiClient creates a new mock instance.
func NewMockapiClient(ctrl *gomock.Controller) *MockapiClient {
	mock := &MockapiClient{ctrl: ctrl}
	mock.recorder = &MockapiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockapiClient) EXPECT() *MockapiClientMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockapiClient) Do(arg0 context.Context, arg1 *http.Request) (*http.Response, []byte, v1.Warnings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(v1.Warnings)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Do indicates an expected call of Do.
func (mr *MockapiClientMockRecorder) Do(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockapiClient)(nil).Do), arg0, arg1)
}

// DoGetFallback mocks base method.
func (m *MockapiClient) DoGetFallback(ctx context.Context, u *url.URL, args url.Values) (*http.Response, []byte, v1.Warnings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoGetFallback", ctx, u, args)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(v1.Warnings)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// DoGetFallback indicates an expected call of DoGetFallback.
func (mr *MockapiClientMockRecorder) DoGetFallback(ctx, u, args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoGetFallback", reflect.TypeOf((*MockapiClient)(nil).DoGetFallback), ctx, u, args)
}

// URL mocks base method.
func (m *MockapiClient) URL(ep string, args map[string]string) *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URL", ep, args)
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// URL indicates an expected call of URL.
func (mr *MockapiClientMockRecorder) URL(ep, args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockapiClient)(nil).URL), ep, args)
}
