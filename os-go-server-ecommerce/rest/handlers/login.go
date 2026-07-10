package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqLoginUser ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLoginUser)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}
	
	user := database.Find(reqLoginUser.Email, reqLoginUser.Password)
	if user == nil {
		http.Error(w, "Inavlid Credentials", http.StatusBadRequest)
		return
	}
	util.SendData(w, user, http.StatusOK)
}
