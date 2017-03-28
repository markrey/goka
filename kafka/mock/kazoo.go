// Automatically generated by MockGen. DO NOT EDIT!
// Source: topic_manager.go

package mock

import (
	kazoo_go "github.com/db7/kazoo-go"
	gomock "github.com/golang/mock/gomock"
)

// Mock of TopicManager interface
type MockTopicManager struct {
	ctrl     *gomock.Controller
	recorder *_MockTopicManagerRecorder
}

// Recorder for MockTopicManager (not exported)
type _MockTopicManagerRecorder struct {
	mock *MockTopicManager
}

func NewMockTopicManager(ctrl *gomock.Controller) *MockTopicManager {
	mock := &MockTopicManager{ctrl: ctrl}
	mock.recorder = &_MockTopicManagerRecorder{mock}
	return mock
}

func (_m *MockTopicManager) EXPECT() *_MockTopicManagerRecorder {
	return _m.recorder
}

func (_m *MockTopicManager) EnsureTableExists(topic string, npar int) error {
	ret := _m.ctrl.Call(_m, "EnsureTableExists", topic, npar)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTopicManagerRecorder) EnsureTableExists(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EnsureTableExists", arg0, arg1)
}

func (_m *MockTopicManager) EnsureStreamExists(topic string, npar int) error {
	ret := _m.ctrl.Call(_m, "EnsureStreamExists", topic, npar)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTopicManagerRecorder) EnsureStreamExists(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EnsureStreamExists", arg0, arg1)
}

func (_m *MockTopicManager) Partitions(topic string) ([]int32, error) {
	ret := _m.ctrl.Call(_m, "Partitions", topic)
	ret0, _ := ret[0].([]int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTopicManagerRecorder) Partitions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Partitions", arg0)
}

func (_m *MockTopicManager) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTopicManagerRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of kzoo interface
type Mockkzoo struct {
	ctrl     *gomock.Controller
	recorder *_MockkzooRecorder
}

// Recorder for Mockkzoo (not exported)
type _MockkzooRecorder struct {
	mock *Mockkzoo
}

func NewMockkzoo(ctrl *gomock.Controller) *Mockkzoo {
	mock := &Mockkzoo{ctrl: ctrl}
	mock.recorder = &_MockkzooRecorder{mock}
	return mock
}

func (_m *Mockkzoo) EXPECT() *_MockkzooRecorder {
	return _m.recorder
}

func (_m *Mockkzoo) Topic(topic string) *kazoo_go.Topic {
	ret := _m.ctrl.Call(_m, "Topic", topic)
	ret0, _ := ret[0].(*kazoo_go.Topic)
	return ret0
}

func (_mr *_MockkzooRecorder) Topic(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Topic", arg0)
}

func (_m *Mockkzoo) Topics() (kazoo_go.TopicList, error) {
	ret := _m.ctrl.Call(_m, "Topics")
	ret0, _ := ret[0].(kazoo_go.TopicList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockkzooRecorder) Topics() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Topics")
}

func (_m *Mockkzoo) CreateTopic(topic string, npar int, rep int, config map[string]string) error {
	ret := _m.ctrl.Call(_m, "CreateTopic", topic, npar, rep, config)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockkzooRecorder) CreateTopic(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateTopic", arg0, arg1, arg2, arg3)
}

func (_m *Mockkzoo) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockkzooRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}