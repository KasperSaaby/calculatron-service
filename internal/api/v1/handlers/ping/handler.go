package ping

import (
	"io"
	"net/http"
)

func Handler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			_, err := io.WriteString(w, "pong\n")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}
