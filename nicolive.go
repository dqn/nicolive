package nicolive

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Nicolive struct {
	client *http.Client
}

func New(mail, password string) (*Nicolive, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Jar: jar}

	resp, err := client.PostForm(
		"https://account.nicovideo.jp/api/v1/login",
		url.Values{
			"mail_tel": {mail},
			"password": {password},
		},
	)
	if err != nil {
		return nil, err
	}

	if f, ok := resp.Header["X-Niconico-Authflag"]; !ok || f[0] != "1" {
		err = fmt.Errorf("failed to login")
		return nil, err
	}

	return &Nicolive{client}, nil
}
