package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var user = User{
	ID:       1,
	Username: "username",
	Password: "password",
	Phone:    "49123454322", //this is a random number
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data Input
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	//compare the user from the request, with the one we defined:
	if user.Username != data.UserName || user.Password != data.PassWord {
		fmt.Fprintln(w, "Please provide valid login details")
		return
	}
	token, err := CreateToken(user.ID)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	outData := Output{
		UserName: data.UserName,
		Token:    token,
	}
	js, err := json.Marshal(outData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}
func CreateToken(userId uint64) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
