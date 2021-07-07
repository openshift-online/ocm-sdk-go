package auth

import (
	"net/http"
)

type SelectFunc = func(r *http.Request) bool

type HandlerSelection struct {
	handler    http.Handler
	selectFunc SelectFunc
}

type AuthSelectorHandler struct {
	selections     []*HandlerSelection
	defaultHandler http.Handler
}

func (h *AuthSelectorHandler) Add(handler http.Handler, selectFunc SelectFunc) {
	h.selections = append(
		h.selections,
		&HandlerSelection{
			handler:    handler,
			selectFunc: selectFunc,
		})
}

func (h *AuthSelectorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, selection := range h.selections {
		if selection.selectFunc(r) {
			selection.handler.ServeHTTP(w, r)
			return
		}
	}
	h.defaultHandler.ServeHTTP(w, r)
}
