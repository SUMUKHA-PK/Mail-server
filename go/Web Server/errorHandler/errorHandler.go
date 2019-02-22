package errorHandler

import "net/http"

func ErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func HttpError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
