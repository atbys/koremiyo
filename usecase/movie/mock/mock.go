// Code generated by MockGen. DO NOT EDIT.
// Source: .\usecase\movie\repository.go

// Package mock_movie is a generated GoMock package.
package mock_movie

import (
	reflect "reflect"

	domain "github.com/atbys/refactor/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockMovieRepository is a mock of MovieRepository interface.
type MockMovieRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMovieRepositoryMockRecorder
}

// MockMovieRepositoryMockRecorder is the mock recorder for MockMovieRepository.
type MockMovieRepositoryMockRecorder struct {
	mock *MockMovieRepository
}

// NewMockMovieRepository creates a new mock instance.
func NewMockMovieRepository(ctrl *gomock.Controller) *MockMovieRepository {
	mock := &MockMovieRepository{ctrl: ctrl}
	mock.recorder = &MockMovieRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieRepository) EXPECT() *MockMovieRepositoryMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockMovieRepository) Find(id int) (*domain.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*domain.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMovieRepositoryMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMovieRepository)(nil).Find), id)
}

// FindClips mocks base method.
func (m *MockMovieRepository) FindClips(fid string) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindClips", fid)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindClips indicates an expected call of FindClips.
func (mr *MockMovieRepositoryMockRecorder) FindClips(fid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindClips", reflect.TypeOf((*MockMovieRepository)(nil).FindClips), fid)
}

// MockMovieCache is a mock of MovieCache interface.
type MockMovieCache struct {
	ctrl     *gomock.Controller
	recorder *MockMovieCacheMockRecorder
}

// MockMovieCacheMockRecorder is the mock recorder for MockMovieCache.
type MockMovieCacheMockRecorder struct {
	mock *MockMovieCache
}

// NewMockMovieCache creates a new mock instance.
func NewMockMovieCache(ctrl *gomock.Controller) *MockMovieCache {
	mock := &MockMovieCache{ctrl: ctrl}
	mock.recorder = &MockMovieCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieCache) EXPECT() *MockMovieCacheMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockMovieCache) Find(arg0 int) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMovieCacheMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMovieCache)(nil).Find), arg0)
}

// Store mocks base method.
func (m *MockMovieCache) Store(arg0 interface{}) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockMovieCacheMockRecorder) Store(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockMovieCache)(nil).Store), arg0)
}