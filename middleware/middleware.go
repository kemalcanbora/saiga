package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"saiga/models"
	"saiga/pkg/helpers"
)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			helpers.HTTPErrorHandler(w, "Token is empty!", http.StatusUnauthorized)
			return
		}

		var secretKey = []byte(os.Getenv("SECRET_KEY"))

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return secretKey, nil
		})

		if err != nil {
			helpers.HTTPErrorHandler(w, "Not Authorized", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "operator" {
				r.Header.Set("role", "operator")
				ctx := context.WithValue(r.Context(), "data", claims)
				handler.ServeHTTP(w, r.WithContext(ctx))
				return
			} else if claims["role"] == "customer" {
				r.Header.Set("role", "customer")
				ctx := context.WithValue(r.Context(), "data", claims)
				handler.ServeHTTP(w, r.WithContext(ctx))
				return
			} else {
				helpers.HTTPErrorHandler(w, "Unauthorized", http.StatusUnauthorized)
			}
		}
	}
}


func CheckUserType(r *http.Request, role string) bool{
	var u models.JwtUserAuth
	user := r.Context().Value("data")
	fmt.Println(user)

	tmp, _ := json.Marshal(user)
	json.Unmarshal(tmp, &u)
	if u.Role != role{
		return false
	}
	return true
}