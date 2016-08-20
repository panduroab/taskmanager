package controllers

import (
	"github.com/panduroab/taskmanager/models"
)

//User Resources
type (
	//For POST - /user/register
	UserResource struct {
		Data models.User `json:"data"`
	}
	//For POST - /user/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}
	//Response for authorized user POST - /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}
	//Model for authentication
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	//Model for authorized user with access token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
)

//Task Resources
type (
	//For POST/PUT - /tasks
	//For GET - /tasks/id
	TaskResource struct {
		Data models.Task `json:"data"`
	}
	//For GET - /tasks
	TasksResource struct {
		Data []models.Task `json:"data"`
	}
)

//Note Resources
type (
	//For POST/PUT - /notes
	//For GET - /notes/id
	NoteResource struct {
		Data NoteModel `json:"data"`
	}
	//For GET - /notes
	NotesResource struct {
		Data []models.TaskNote `json:"data"`
	}
	//Model for a TaskNote
	NoteModel struct {
		TaskID      string `json:"taskid"`
		Description string `json:"description"`
	}
)
