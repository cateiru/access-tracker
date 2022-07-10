package base

import (
	"errors"
	"net/http"
)

type Base struct {
	w        http.ResponseWriter
	r        *http.Request
	UserData *UserData
}

func New(w http.ResponseWriter, r *http.Request) *Base {
	return &Base{
		w:        w,
		r:        r,
		UserData: ParseUA(r),
	}
}

func (c Base) GetIPAddress() string {
	ip := c.r.Header.Get("x-forwarded-for")
	if ip != "" {
		return ip
	}
	return c.r.RemoteAddr
}

func (c Base) GetQuery(key string) (string, error) {
	query := c.r.URL.Query().Get(key)

	if len(query) == 0 {
		return "", errors.New("query is empty")
	}

	return query, nil
}
