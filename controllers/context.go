package controllers

import (
	"github.com/panduroab/taskmanager/common"
	"gopkg.in/mgo.v2"
)

//Context used for maintaining HTTP Request context
type Context struct {
	MongoSession *mgo.Session
}

//Close mgo.Session
func (c *Context) Close() {
	c.MongoSession.Close()
}

//DbCollection Return mgo.collection for the given name
func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.Database).C(name)
}

//NewContext Create a new Context for each HTTP Request
func NewContext() *Context {
	session := common.GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}
