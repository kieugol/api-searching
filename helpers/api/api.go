package api

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type HttClientI interface {
	Get(url string, timeout int, header map[string]string) ([]byte, int)
}

type HttClient struct {
}

func (httpReq *HttClient) Get(url string, timeout int, header map[string]string) ([]byte, int) {
	httpClient := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		logrus.WithField("api", url).WithField("error", err.Error()).Error("HTTP_REQUEST_ERROR")
		return nil, http.StatusInternalServerError
	}

	req.Header.Add("accept", "application/json")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	res, err := httpClient.Do(req)
	if err != nil {
		logrus.WithField("api", url).WithField("error", err.Error()).Error("HTTP_CALLING_ERROR")
		return nil, http.StatusInternalServerError
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.WithField("api", url).WithField("error", string(body)).Error("READ_BODY_ERROR")
		return nil, http.StatusInternalServerError
	}

	logrus.WithField("api", url).WithField("response", string(body)).Info("API_RETURN_DATA")

	return body, res.StatusCode
}
