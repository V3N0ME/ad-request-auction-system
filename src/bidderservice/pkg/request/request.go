package request

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const defaultTimeout = 25
const defaultMaxOpenConnections = 100

//Config holds the configuration details of the middleware
type Config struct {
	Timeout            time.Duration
	MaxOpenConnections int
}

//CustomHTTP default
type CustomHTTP struct {
	config    Config
	client    http.Client
	httpQueue chan struct{}
}

//Request is the model of the request to make
type Request struct {
	URL     string
	Method  string
	Payload []byte
	Headers map[string]string
}

//New returns a new instance of request
func New(config Config) *CustomHTTP {
	if config.Timeout == 0 {
		config.Timeout = time.Second * defaultTimeout
	}

	if config.MaxOpenConnections == 0 {
		config.MaxOpenConnections = defaultMaxOpenConnections
	}

	return &CustomHTTP{
		config:    config,
		httpQueue: make(chan struct{}, config.MaxOpenConnections),
		client: http.Client{
			Transport: &http.Transport{
				ResponseHeaderTimeout: config.Timeout,
			},
		},
	}
}

//MakeRequest makes http requests
func (c *CustomHTTP) MakeRequest(request Request) (string, int, error) {

	didRetry := false

	//waits until queue has space left
	c.httpQueue <- struct{}{}
	defer func() {
		//to avoid queue from being occupied until request is succeeded after retry
		//instead queue is poped if an error occurs before retrying
		if !didRetry {
			<-c.httpQueue
		}
	}()

	req, err := http.NewRequest(request.Method, request.URL, bytes.NewBuffer(request.Payload))

	for v, k := range request.Headers {
		req.Header.Set(v, k)
	}

	if err != nil {
		return "", 0, err
	}

	resp, err := c.client.Do(req)
	//close the connection after the transaction is complete
	//to avoid keeping the file descriptor open to be reused for the next connection
	req.Close = true

	if err != nil {
		return "", 0, err
	}

	bytBody, err := ioutil.ReadAll(resp.Body)
	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	strBody := string(bytBody)

	if resp.StatusCode > 423 {

		return strBody, resp.StatusCode, errors.New("Internal Server Error")
	}

	return strBody, resp.StatusCode, nil
}
