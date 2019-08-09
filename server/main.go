package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	key = []byte("mysecretkey")
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is secret")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}

				return key, nil

			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "not authorized")
		}

	})
}

func HandleRequest() {
	http.Handle("/secret", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("simple server")
	HandleRequest()
}
