package sap_api_post_header_setup

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	"time"

	"golang.org/x/xerrors"
)

type SAPPostClient struct {
	jar                   http.CookieJar
	csrfToken             string
	refreshTokenURL       string
	user                  string
	pass                  string
	retryMaxCnt           int
	retryIntervalMilliSec int
}

type SAPPostClientOption interface {
	User() string
	Pass() string
	RefreshTokenURL() string
	RetryMax() int
	RetryInterval() int
}

func NewSAPPostClientWithOption(o SAPPostClientOption) *SAPPostClient {
	return NewSAPPostClient(
		o.User(), o.Pass(), o.RefreshTokenURL(), o.RetryMax(), o.RetryInterval(),
	)
}
func NewSAPPostClient(sapUserName, sapPass, refreshTokenURL string, retryMaxCnt, retryIntervalMilliSec int) *SAPPostClient {
	if refreshTokenURL == "" {
		refreshTokenURL = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	}
	j, _ := cookiejar.New(nil)
	c := &SAPPostClient{
		jar:                   j,
		csrfToken:             "",
		refreshTokenURL:       refreshTokenURL,
		user:                  sapUserName,
		pass:                  sapPass,
		retryMaxCnt:           retryMaxCnt,
		retryIntervalMilliSec: retryIntervalMilliSec,
	}
	return c
}

func (c *SAPPostClient) POST(url string, params map[string]string, body string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return nil, xerrors.Errorf("cannot create request: %w", err)
	}
	c.setHeader(req, url, body)
	c.setParams(req, params)
	client := &http.Client{
		Jar: c.jar,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, xerrors.Errorf("request returns error: %w", err)
	}
	if res.StatusCode == http.StatusUnauthorized ||
		res.StatusCode == http.StatusForbidden {
		c.updateToken()
		req, _ := http.NewRequest("POST", url, strings.NewReader(body))
		c.setHeader(req, url, body)
		c.setParams(req, params)
		res, err = client.Do(req)
	}
	if err != nil {
		return nil, xerrors.Errorf("request returns error: %w", err)
	}
	return res, nil
}

func (c *SAPPostClient) setHeader(req *http.Request, url string, body string) {
	req.SetBasicAuth(c.user, c.pass)
	req.Header.Add("x-csrf-token", c.csrfToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
}

func (c *SAPPostClient) setParams(req *http.Request, params map[string]string) {
	parameter := req.URL.Query()
	for k, v := range params {
		parameter.Add(k, v)
	}
	req.URL.RawQuery = parameter.Encode()
}

func (c *SAPPostClient) updateToken() {
	req, _ := http.NewRequest("GET", c.refreshTokenURL, nil)
	req.Header.Add("x-csrf-token", "Fetch")
	req.SetBasicAuth(c.user, c.pass)
	var err error
	// do while 実装のために評価はtrue
	for cnt := 1; true; cnt++ {
		res, err := (&http.Client{
			Jar: c.jar,
		}).Do(req)
		if err == nil && res != nil && res.Header.Get("x-csrf-token") != "" {
			c.csrfToken = res.Header.Get("x-csrf-token")
			return
		}
		// do while 実装のための評価
		if cnt >= c.retryMaxCnt {
			break
		}
		time.Sleep(time.Duration(c.retryIntervalMilliSec) * time.Millisecond)
	}
	// 最後のエラーだけ拾う
	if err != nil {
		fmt.Fprintf(os.Stderr, "token update error: %+v", err)
	}
}
