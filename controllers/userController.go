package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/panduroab/taskmanager/common"
	"github.com/panduroab/taskmanager/data"
	"github.com/panduroab/taskmanager/models"
)

//Register a new User document
func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	//Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			500,
		)
		return
	}
	user := &dataResource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}
	//Insert User document
	repo.CreateUser(user)
	//Clean-up the hashpassword to eliminate it from response
	user.HashPassword = nil
	j, err := json.Marshal(UserResource{Data: *user})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//Login user into the system
func Login(w http.ResponseWriter, r *http.Request) {
	var dataResource LoginResource
	var token string
	//Decode the incoming login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Login data",
			500,
		)
		return
	}
	loginModel := dataResource.Data
	loginUser := models.User{
		Email:    loginModel.Email,
		Password: loginModel.Password,
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}
	//Authenticate the login User
	user, err := repo.Login(loginUser)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Login credentials",
			401,
		)
		return
	}
	//Generate JWT token
	token, err = common.GenerateJWT(user.Email, "member")
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Error while generating the access token",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	user.HashPassword = nil
	authUser := AuthUserModel{
		User:  user,
		Token: token,
	}
	j, err := json.Marshal(AuthUserResource{Data: authUser})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
