package baidu

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
)

type Baidu struct {
	AppId     string
	ApiKey    string
	ApiSecret string

	tkn     *AccessToken
	baseURL *url.URL
}

// NewBaidu
func NewBaidu(apiKey, apiSecret string) *Baidu {
	b := &Baidu{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	}
	b.baseURL, _ = url.Parse("https://aip.baidubce.com/")
	return b
}

// path
func (b *Baidu) path(path string) *url.URL {
	u, _ := url.Parse(b.baseURL.String())
	u.Path = path
	return u
}

type AccessToken struct {
	AccessToken   string `json:"access_token"`
	SessionKey    string `json:"session_key"`
	Scope         string `json:"scope"`
	RefreshToken  string `json:"refresh_token"`
	SessionSecret string `json:"session_secret"`
	ExpiresIn     int64  `json:"expires_in"`
}

// AccessToken
func (b *Baidu) GetAccessToken() (*AccessToken, error) {
	if b.tkn != nil {
		return b.tkn, nil
	}
	u := b.path("/oauth/2.0/token")
	q := u.Query()
	q.Set("grant_type", "client_credentials")
	q.Set("client_id", b.ApiKey)
	q.Set("client_secret", b.ApiSecret)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		logrus.Errorf("got access token failed, %s", err)
		return nil, err
	}

	tkn := &AccessToken{}
	err = json.NewDecoder(resp.Body).Decode(tkn)
	if err != nil {
		logrus.Errorf("decode access token json failed, %s", err)
		return nil, err
	}

	b.tkn = tkn
	return tkn, nil
}

type BaiduResponse struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	LogId     int64  `json:"log_id"`
	Timestamp int64  `json:"timestamp"`
}
