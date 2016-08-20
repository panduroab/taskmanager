package controllers

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"github.com/panduroab/taskmanager/common"
	"github.com/panduroab/taskmanager/data"
	"github.com/panduroab/taskmanager/models"
)

//CreateNote func
func CreateNote(w http.ResponseWriter, r *http.Request) {
	var noteResource NoteResource
	err := json.NewDecoder(r.Body).Decode(&noteResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Note data",
			500,
		)
		return
	}
	noteModel := &noteResource.Data
	note := &models.TaskNote{
		TaskID:      bson.ObjectIdHex(noteModel.TaskID),
		Description: noteModel.Description,
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("notes")
	repo := &data.NoteRepository{c}
	repo.Create(note)
	j, err := json.Marshal(note)
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
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//UpdateNote func
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	//Get the id from the url
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource NoteResource
	//Decode the incoming json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Note data",
			500,
		)
		return
	}
	noteModel := dataResource.Data
	note := &models.TaskNote{
		ID:          id,
		Description: noteModel.Description,
	}
	//Update the TaskNote
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("notes")
	repo := &data.NoteRepository{c}
	if err := repo.Update(note); err != nil {
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

//GetNotes func
func GetNotes(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("notes")
	repo := &data.NoteRepository{c}
	notes := repo.GetAll()
	j, err := json.Marshal(NotesResource{Data: notes})
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

//GetNoteByID func
func GetNoteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("notes")
	repo := &data.NoteRepository{c}
	note, err := repo.GetByID(id)
	if err != nil {
		if err := mgo.ErrNotFound; err != nil {
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
	j, err := json.Marshal(note)
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

//GetNotesByTask func
func GetNotesByTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("notes")
	repo := &data.NoteRepository{c}
	notes := repo.GetByTask(id)
	j, err := json.Marshal(NotesResource{Data: notes})
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

//DeleteNote func
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	//Get the id from the URL
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("notes")
	repo := &data.NoteRepository{c}
	//Delete the TaskNote
	if err := repo.Delete(id); err != nil {
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
