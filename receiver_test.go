package main

import (
	"errors"
	"testing"

	"github.com/containrrr/shoutrrr/pkg/types"
	"github.com/mattbun/smtprrr/internal/notify"
	"github.com/mattbun/smtprrr/internal/smtp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSender struct {
	mock.Mock
}

func (m *MockSender) Send(message string, params *types.Params) []error {
	args := m.MethodCalled("Send", message, params)
	return args.Get(0).([]error)
}

func TestNewNotifierMessageReceiver_Success(t *testing.T) {
	_, err := NewNotifierMessageReceiver([]string{"pushover://shoutrrr:token@user/"})
	assert.NoError(t, err)
}

func TestNewNotifierMessageReceiver_Failure(t *testing.T) {
	_, err := NewNotifierMessageReceiver([]string{"invalid-url"})
	assert.Error(t, err)
}

func TestOnMessage_Success(t *testing.T) {
	mockSender := MockSender{}
	mockNotifier := notify.Notifier{
		Sender: &mockSender,
	}
	receiver := NotifierMessageReceiver{
		Notifier: mockNotifier,
	}

	mockSender.On("Send", mock.Anything, mock.Anything).Return([]error{})

	err := receiver.OnMessage(smtp.Message{
		Subject: "some subject",
		Body:    "some body",
	})

	assert.NoError(t, err)
}

func TestOnMessage_Failure(t *testing.T) {
	mockSender := MockSender{}
	mockNotifier := notify.Notifier{
		Sender: &mockSender,
	}
	receiver := NotifierMessageReceiver{
		Notifier: mockNotifier,
	}

	mockSender.On("Send", mock.Anything, mock.Anything).Return([]error{errors.New("some error")})

	err := receiver.OnMessage(smtp.Message{
		Subject: "some subject",
		Body:    "some body",
	})

	assert.Error(t, err)
}
