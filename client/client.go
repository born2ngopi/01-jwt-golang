package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	key = []byte("mysecretkey")
)

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := CreateJWTToken()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, validToken)
}

func CreateJWTToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "needkopi"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Errorf("error in %s\n", err.Error())
		return "", nil
	}

	return tokenString, nil
}

func HandleRequest() {
	http.HandleFunc("/token", homePage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("server running")
	HandleRequest()
}
