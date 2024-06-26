// Code generated by MockGen. DO NOT EDIT.
// Source: ./article_tag.go
//
// Generated by this command:
//
//	mockgen -source ./article_tag.go -destination ./mock_repo/article_tag.go -package mock_repo
//
// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	context "context"
	reflect "reflect"

	model "github.com/sunjin110/folio/golio/domain/model"
	repository "github.com/sunjin110/folio/golio/domain/repository"
	gomock "go.uber.org/mock/gomock"
)

// MockArticleTag is a mock of ArticleTag interface.
type MockArticleTag struct {
	ctrl     *gomock.Controller
	recorder *MockArticleTagMockRecorder
}

// MockArticleTagMockRecorder is the mock recorder for MockArticleTag.
type MockArticleTagMockRecorder struct {
	mock *MockArticleTag
}

// NewMockArticleTag creates a new mock instance.
func NewMockArticleTag(ctrl *gomock.Controller) *MockArticleTag {
	mock := &MockArticleTag{ctrl: ctrl}
	mock.recorder = &MockArticleTagMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleTag) EXPECT() *MockArticleTagMockRecorder {
	return m.recorder
}

// CountTotal mocks base method.
func (m *MockArticleTag) CountTotal(ctx context.Context, search *repository.ArticleTagSearch) (int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountTotal", ctx, search)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountTotal indicates an expected call of CountTotal.
func (mr *MockArticleTagMockRecorder) CountTotal(ctx, search any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountTotal", reflect.TypeOf((*MockArticleTag)(nil).CountTotal), ctx, search)
}

// Delete mocks base method.
func (m *MockArticleTag) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockArticleTagMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockArticleTag)(nil).Delete), ctx, id)
}

// Find mocks base method.
func (m *MockArticleTag) Find(ctx context.Context, sortType repository.SortType, paging *repository.Paging, search *repository.ArticleTagSearch) ([]*model.ArticleTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, sortType, paging, search)
	ret0, _ := ret[0].([]*model.ArticleTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockArticleTagMockRecorder) Find(ctx, sortType, paging, search any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockArticleTag)(nil).Find), ctx, sortType, paging, search)
}

// FindByIDs mocks base method.
func (m *MockArticleTag) FindByIDs(ctx context.Context, ids []string) ([]*model.ArticleTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIDs", ctx, ids)
	ret0, _ := ret[0].([]*model.ArticleTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIDs indicates an expected call of FindByIDs.
func (mr *MockArticleTagMockRecorder) FindByIDs(ctx, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIDs", reflect.TypeOf((*MockArticleTag)(nil).FindByIDs), ctx, ids)
}

// Get mocks base method.
func (m *MockArticleTag) Get(ctx context.Context, id string) (*model.ArticleTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*model.ArticleTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockArticleTagMockRecorder) Get(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockArticleTag)(nil).Get), ctx, id)
}

// Insert mocks base method.
func (m *MockArticleTag) Insert(ctx context.Context, tag *model.ArticleTag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, tag)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockArticleTagMockRecorder) Insert(ctx, tag any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockArticleTag)(nil).Insert), ctx, tag)
}

// Update mocks base method.
func (m *MockArticleTag) Update(ctx context.Context, tag *model.ArticleTag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, tag)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockArticleTagMockRecorder) Update(ctx, tag any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArticleTag)(nil).Update), ctx, tag)
}
