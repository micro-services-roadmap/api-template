package middleware

import (
	"net/http"
)

type CheckUrlMiddleware struct {
}

func NewCheckUrlMiddleware() *CheckUrlMiddleware {
	return &CheckUrlMiddleware{}
}

func (m *CheckUrlMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}
