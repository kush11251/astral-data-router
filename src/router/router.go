package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func StartRouter() {
	http.HandleFunc("/data", handleData)
	http.ListenAndServe(":8080", nil)
}

func handleData(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	_ = json.NewDecoder(r.Body).Decode(&data)
	fmt.Println(data)
	w.WriteHeader(http.StatusOK)
}