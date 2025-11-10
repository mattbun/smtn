package smtp

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var rawMessage string = `From: from@test.com
Date: Sun, 09 Nov 2025 21:40:14 -0500
Message-ID: <65184fb2838c074f74d364e6a5d9b2d0@test.com>
Subject:Hey now

Some body once told me`

type MockMessageReceiver struct {
	mock.Mock
}

func (r *MockMessageReceiver) OnMessage(message Message) error {
	args := r.MethodCalled("OnMessage", message)
	return args.Error(0)
}

func TestData_Success(t *testing.T) {
	mockMessageReceiver := MockMessageReceiver{}
	mockMessageReceiver.On("OnMessage", mock.Anything).Return(nil)
	session := Session{MessageReceiver: &mockMessageReceiver}

	err := session.Data(strings.NewReader(rawMessage))

	assert.NoError(t, err)
	mockMessageReceiver.AssertCalled(t, "OnMessage", Message{
		Subject: "Hey now",
		Body:    "Some body once told me",
	})
}

func TestData_Success_TrimWhitespace(t *testing.T) {
	mockMessageReceiver := MockMessageReceiver{}
	mockMessageReceiver.On("OnMessage", mock.Anything).Return(nil)
	session := Session{MessageReceiver: &mockMessageReceiver}

	err := session.Data(strings.NewReader(rawMessage + "     "))

	assert.NoError(t, err)
	mockMessageReceiver.AssertCalled(t, "OnMessage", Message{
		Subject: "Hey now",
		Body:    "Some body once told me",
	})
}

func TestData_Failure_OnMessage(t *testing.T) {
	mockErr := errors.New("some error")
	mockMessageReceiver := MockMessageReceiver{}
	mockMessageReceiver.On("OnMessage", mock.Anything).Return(mockErr)
	session := Session{MessageReceiver: &mockMessageReceiver}

	err := session.Data(strings.NewReader(rawMessage))

	assert.Error(t, err)
}

func TestData_Failure_ReadMessage(t *testing.T) {
	mockMessageReceiver := MockMessageReceiver{}
	mockMessageReceiver.On("OnMessage", mock.Anything).Return(nil)
	session := Session{MessageReceiver: &mockMessageReceiver}

	err := session.Data(strings.NewReader(""))

	assert.Error(t, err)
}
