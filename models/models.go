package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirtsName   string         `json:"firstName"`
		LastName    string         `json:"lastName"`
		Email       string         `json:"email"`
		Password    string         `json:"password,omitempty"`
		HashPassword []byte         `json:"hashPasswd,omitempty"`
	}
	Task struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy   string        `json:"createdby"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Due         time.Time     `json:"due,omitempty"`
		Status      string        `json:"status,omitempty"`
		Tags        []string      `json:"tags,omitempty"`
	}
		TaskNote struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		TaskId      bson.ObjectId `json:"taskid"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
	}
)