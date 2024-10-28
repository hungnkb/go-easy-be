package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginDto struct {
	Username string `string;json:"username"`
	Password string `string;json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var body LoginDto
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
}
