package httputil

import (
	"io/ioutil"
	"net/http"

	"github.com/ryantking/jetbugs/mylib/pkg"
)

func Request(c pkg.Config) (int, string, error) {
	req, err := http.NewRequest(c.Method, c.URL, c.Body)
	if err != nil {
		return 0, "", err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, "", err
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, "", err
	}

	return res.StatusCode, string(b), nil
}
