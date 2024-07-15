package controllers

import (
	"gopkg.in/mgo.v2"

	"github.com/mahdi-eth/go-taskmanager/common"
)

type Context struct {
	MongoSession *mgo.Session
	User         string
}

func (c *Context) Close() {
	c.MongoSession.Close()
}

func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.Database).C(name)
}

func NewContext() *Context {
	session := common.GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}