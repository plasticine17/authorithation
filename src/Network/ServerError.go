package Network

import "net/http"

func setServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - InternalServerError"))
}
