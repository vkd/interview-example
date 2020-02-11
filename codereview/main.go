package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/calculate", handler)
	http.ListenAndServe(":5000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ns := r.URL.Query().Get("n")
	n, _ := strconv.Atoi(ns)
	res := calculate(n)
	response, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func calculate(n int) int {
	if n == 1 {
		return n
	}
	return n * calculate(n-1)
}
