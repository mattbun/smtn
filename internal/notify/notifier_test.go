package notify

import (
	"errors"
	"testing"

	"github.com/containrrr/shoutrrr/pkg/types"
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

func TestNotify_Success(t *testing.T) {
	mockSender := MockSender{}
	mockNotifier := Notifier{
		Sender: &mockSender,
	}
	mockSender.On("Send", mock.Anything, mock.Anything).Return([]error{})

	err := mockNotifier.Notify(NotifyInput{
		Title: "some title",
		Body:  "some body",
	})

	assert.NoError(t, err)

	mockSender.AssertNumberOfCalls(t, "Send", 1)
	args := mockSender.Calls[0].Arguments
	assert.Equal(t, "some body", args[0])

	params, isParams := args[1].(*types.Params)
	assert.True(t, isParams, "Should be types.Params")
	title, found := params.Title()
	assert.True(t, found)
	assert.Equal(t, "some title", title)
}

func TestNotify_Failure(t *testing.T) {
	mockSender := MockSender{}
	mockNotifier := Notifier{
		Sender: &mockSender,
	}
	mockSender.On("Send", mock.Anything, mock.Anything).Return([]error{errors.New("some error")})

	err := mockNotifier.Notify(NotifyInput{
		Title: "some title",
		Body:  "some body",
	})

	assert.Error(t, err)
}

func TestNewNotifier_Success(t *testing.T) {
	_, err := NewNotifier([]string{"pushover://shoutrrr:token@user/"})
	assert.NoError(t, err)
}

func TestNewNotifier_Failure(t *testing.T) {
	_, err := NewNotifier([]string{"bad-url"})
	assert.Error(t, err)
}
