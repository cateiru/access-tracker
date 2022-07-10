package base

import "net/http"

type Wrapper struct {
	mux *http.ServeMux
}

func NewWrapper(mux *http.ServeMux) *Wrapper {
	return &Wrapper{
		mux: mux,
	}
}

func (c Wrapper) HandleFunc(pattern string, handler func(*Base)) {
	h := func(w http.ResponseWriter, r *http.Request) {
		base := New(w, r)
		handler(base)
	}

	c.mux.HandleFunc(pattern, h)
}
