package controllers

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"github.com/panduroab/taskmanager/common"
	"github.com/panduroab/taskmanager/data"
)

//CreateTask Insert a new task document
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var dataResources TaskResource
	//Decode the incoming Task json
	err := json.NewDecoder(r.Body).Decode(&dataResources)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Task data",
			500,
		)
		return
	}
	task := &dataResources.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}
	//Insert a task document
	repo.Create(task)
	j, err := json.Marshal(TaskResource{Data: *task})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has ocurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//GetTasks Return all Task documents
func GetTasks(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}
	tasks := repo.GetAll()
	j, err := json.Marshal(TasksResource{Data: tasks})
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
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//GetTaskByID Return a requested task by id
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	//Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}
	task, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	j, err := json.Marshal(task)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetTasksByUser Return all Tasks created by a User
func GetTasksByUser(w http.ResponseWriter, r *http.Request) {
	//Get id from the incoming url
	vars := mux.Vars(r)
	user := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}
	tasks := repo.GetByUser(user)
	j, err := json.Marshal(TasksResource{Data: tasks})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//UpdateTask an existing Task document
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	//Get id from the incoming url
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource TaskResource
	//Decode the incoming Task json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Task data",
			500,
		)
		return
	}
	task := &dataResource.Data
	task.ID = id
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}
	//Update an existing Task document
	if err := repo.Update(task); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//DeleteTask an existing Task document
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}
	//Delete an existing Task document
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
