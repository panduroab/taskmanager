package data

import (
	"time"

	"github.com/panduroab/taskmanager/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//NoteRepository struct
type NoteRepository struct {
	C *mgo.Collection
}

//Create inserts a new Note
func (r *NoteRepository) Create(note *models.TaskNote) error {
	note.ID = bson.NewObjectId()
	note.CreatedOn = time.Now()
	err := r.C.Insert(&note)
	return err
}

//Update function updates a TaskNote
func (r *NoteRepository) Update(note *models.TaskNote) error {
	err := r.C.Update(bson.M{"_id": note.ID},
		bson.M{"$set": bson.M{
			"description": note.Description,
		}})
	return err
}

//Delete function deletes a TaskNote
func (r *NoteRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//GetByTask return all the Notes form a Task
func (r *NoteRepository) GetByTask(id string) []models.TaskNote {
	var notes []models.TaskNote
	taskid := bson.ObjectIdHex(id)
	iter := r.C.Find(bson.M{"taskid": taskid}).Iter()
	result := models.TaskNote{}
	for iter.Next(&result) {
		notes = append(notes, result)
	}
	return notes
}

//GetAll return all the TaskNote elements
func (r *NoteRepository) GetAll() []models.TaskNote {
	var notes []models.TaskNote
	iter := r.C.Find(nil).Iter()
	result := models.TaskNote{}
	for iter.Next(&result) {
		notes = append(notes, result)
	}
	return notes
}

//GetByID return a TaskNote model
func (r *NoteRepository) GetByID(id string) (note models.TaskNote, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&note)
	return
}
