package handlers

import "net/http"

type Handler struct {
	h func(http.ResponseWriter, *http.Request)
}
