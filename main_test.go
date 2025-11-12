package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/smtp"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	// smtpPort is the port that the test SMTP server will listen on.
	smtpPort = 10025

	// rawMessage is the email message that will be sent to the test SMTP server.
	rawMessage = `From:blah@blah
Subject:Hey now

Some body once told me`
)

// ShoutrrrPayload is the 'json' template payload sent to generic webhooks with shoutrrr.
type ShoutrrrPayload struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

func TestE2E(t *testing.T) {
	ctx := context.Background()

	// Start HTTP server
	httpEndpointCalled := false
	var httpBody []byte
	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpEndpointCalled = true
		httpBody, _ = io.ReadAll(r.Body)
		fmt.Fprintf(w, "success")
	}))
	defer httpServer.Close()

	// Parse HTTP server url so we can get the host and port out of it
	httpUrl, err := url.Parse(httpServer.URL)
	assert.NoError(t, err)

	// Start smtn and tell it to hit the http server
	go func() {
		args := []string{
			"smtn",
			"--notification-url",
			fmt.Sprintf("generic://%s:%s?template=json&disabletls=yes", httpUrl.Hostname(), httpUrl.Port()),
			"--port",
			fmt.Sprintf("%d", smtpPort),
			"--allow-insecure",
		}
		cmd.Run(ctx, args)
	}()

	// Wait for SMTP server to start
	time.Sleep(1 * time.Second)

	// Send mail to SMTP server
	err = smtp.SendMail(
		fmt.Sprintf("localhost:%d", smtpPort),
		nil, // No auth
		"from@localhost",
		[]string{"to@localhost"},
		[]byte(rawMessage),
	)
	assert.NoError(t, err)

	// Verify HTTP server received a request and that the body is not empty
	assert.Eventually(t, func() bool {
		return httpEndpointCalled && len(httpBody) > 0
	}, 5*time.Second, 10*time.Millisecond)

	// Check that the request body is the same as the mail subject and body
	parsedPayload := ShoutrrrPayload{}
	err = json.Unmarshal(httpBody, &parsedPayload)
	assert.NoError(t, err)
	assert.Equal(t, "Hey now", parsedPayload.Title)
	assert.Equal(t, "Some body once told me", parsedPayload.Message)
}
