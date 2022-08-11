package internal

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func WithLogger(client *http.Client, logger *logrus.Logger) (result *http.Client) {
	result = client
	result.Transport = LoggerTransport{
		logger: logger,
		next:   client.Transport,
	}

	return result
}

type LoggerTransport struct {
	logger *logrus.Logger
	next   http.RoundTripper
}

func (t LoggerTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	t.logger.Debugf("requesting... %s %s", request.Method, request.URL)

	response, err := t.next.RoundTrip(request)
	if err != nil {
		t.logger.Errorf("ERROR %s", response.Status)
	} else {
		t.logger.Debugf("%s %d", response.Status, response.ContentLength)
	}

	return response, err
}
