package api

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type IHttpClient interface {
	SendGet(agrs Params) (string, int)
}

type HttpClient struct {
}
type Params struct {
	URL     string
	Timeout int
	Header  map[string]string
}

func NewHttClient() *HttpClient {
	return &HttpClient{}
}

func (httpReq *HttpClient) SendGet(agrs Params) (string, int) {
	httpClient := &http.Client{
		Timeout: time.Duration(agrs.Timeout) * time.Second,
	}
	req, err := http.NewRequest("GET", agrs.URL, nil)

	if err != nil {
		logrus.WithField("api", agrs.URL).WithField("error", err.Error()).Error("HTTP_REQUEST_ERROR")
		return "", http.StatusInternalServerError
	}

	req.Header.Add("accept", "application/json")
	for k, v := range agrs.Header {
		req.Header.Set(k, v)
	}
	res, err := httpClient.Do(req)
	if err != nil {
		logrus.WithField("api", agrs.URL).WithField("error", err.Error()).Error("HTTP_CALLING_ERROR")
		return "", http.StatusInternalServerError
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.WithField("api", agrs.URL).WithField("error", err.Error()).Error("READ_BODY_ERROR")
		return "", http.StatusInternalServerError
	}
	strBody := string(body)

	logrus.WithField("api", agrs.URL).WithField("response", strBody).Info("API_RETURN_DATA")

	return strBody, res.StatusCode
}
