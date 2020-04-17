package v2

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func request(params map[string]string, method string) (*http.Request, error) {
	host := fmt.Sprintf("https://%s", server())
	u, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}
	for k, v := range params {
		q.Add(k, v)
	}

	u.RawQuery = q.Encode()
	return http.NewRequest(method, u.String(), nil)
}

func server() string {
	return "ccs.api.qcloud.com/v2/index.php"
}

func sign(secret, method string, params map[string]string) string {
	data := method + server()
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var qs []string
	for _, k := range keys {
		qs = append(qs, fmt.Sprintf("%s=%s", k, params[k]))
	}
	data = data + "?" + strings.Join(qs, "&")

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func doRequest(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if int(resp.StatusCode/100) != 2 {
		return nil, fmt.Errorf("http status not 2xx: %d %s", resp.StatusCode, string(body))
	}

	return body, nil
}
