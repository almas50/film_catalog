package handler

import (
	"net/http"
)

func (h *Handler) middlewareAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message": "No basic auth present"}`))
			return
		}
		user, err := h.Service.GetUser(username, password)
		if err != nil {
			w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message": "Invalid username or password"}`))
			return
		}
		method := r.Method
		if method == "POST" || method == "UPDATE" || method == "DELETE" {
			if user.Role == "user" {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte(`{"message": "You don’t have permission "}`))
				return
			}
		}
		next.ServeHTTP(w, r)
	}
}
