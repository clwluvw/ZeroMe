package promquerier

import "net/http"

type headerTransport struct {
	headers map[string]string
}

func (t *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.headers != nil {
		for k, v := range t.headers {
			req.Header.Add(k, v)
		}
	}

	return http.DefaultTransport.RoundTrip(req)
}
