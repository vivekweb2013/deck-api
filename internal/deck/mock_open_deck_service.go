// Code generated by MockGen. DO NOT EDIT.
// Source: open_deck_service.go

// Package deck is a generated GoMock package.
package deck

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOpenDeckService is a mock of OpenDeckService interface.
type MockOpenDeckService struct {
	ctrl     *gomock.Controller
	recorder *MockOpenDeckServiceMockRecorder
}

// MockOpenDeckServiceMockRecorder is the mock recorder for MockOpenDeckService.
type MockOpenDeckServiceMockRecorder struct {
	mock *MockOpenDeckService
}

// NewMockOpenDeckService creates a new mock instance.
func NewMockOpenDeckService(ctrl *gomock.Controller) *MockOpenDeckService {
	mock := &MockOpenDeckService{ctrl: ctrl}
	mock.recorder = &MockOpenDeckServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenDeckService) EXPECT() *MockOpenDeckServiceMockRecorder {
	return m.recorder
}

// OpenDeck mocks base method.
func (m *MockOpenDeckService) OpenDeck(deckID string) (*Deck, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenDeck", deckID)
	ret0, _ := ret[0].(*Deck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenDeck indicates an expected call of OpenDeck.
func (mr *MockOpenDeckServiceMockRecorder) OpenDeck(deckID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenDeck", reflect.TypeOf((*MockOpenDeckService)(nil).OpenDeck), deckID)
}