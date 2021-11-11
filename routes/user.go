package routes

import (
	"encoding/json"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"saiga/models"
	"saiga/pkg/clients"
	"saiga/pkg/helpers"
	"time"
)

func SignUp(response http.ResponseWriter, request *http.Request) {
	var user models.User

	response.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(request.Body).Decode(&user)
	validate := validator.New()
	errVal := validate.Struct(user)

	if errVal != nil {
		helpers.HTTPErrorHandler(response, "Email or Password field is empty!", http.StatusUnauthorized)
		return
	}

	userCheck := clients.Mongo.FindUserWithEmail(user)
	if userCheck.Email == user.Email {
		helpers.HTTPErrorHandler(response, "This email is already registered!", http.StatusNotAcceptable)
		return
	}

	if err != nil {
		helpers.HTTPErrorHandler(response, `{"message":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	user.Password = helpers.GetHash([]byte(user.Password))
	user.CreatedTime = time.Now().Unix()
	if user.Role == "" {
		user.Role = "customer"
	}

	result, _ := clients.Mongo.Insert(user, "user")
	json.NewEncoder(response).Encode(result)
}

func UserLogin(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(request.Body).Decode(&user)
	validate := validator.New()
	errVal := validate.Struct(user)

	if errVal != nil {
		helpers.HTTPErrorHandler(response, "Email or Password field is empty!", http.StatusUnauthorized)
		return
	}

	result := clients.Mongo.FindUserWithEmail(user)
	passErr := helpers.CheckPasswordHash(user.Password, result.Password)

	if passErr != true {
		log.Println(passErr)
		response.Write([]byte(`{"response":"Wrong Password!"}`))
		return
	}

	jwtToken, err := helpers.GenerateJWT(result)
	if err != nil {
		helpers.HTTPErrorHandler(response, `{"message":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}
	response.Write([]byte(`{"token":"` + jwtToken + `"}`))
}
