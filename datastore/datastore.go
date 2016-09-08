package datastore

import (
	"log"

	"github.com/cream-lab/vanilla/model"
)

type DataStore interface {
	Name() string
	GetUserCount() int
	GetUserBySeq(seq int) model.User
	AddUser(user model.User) (int, int)
}

type InMemoryDataStore struct {
	Users map[int]model.User
}

func (datastore *InMemoryDataStore) Name() string {
	return "InMemoryDatastore"
}

func (datastore *InMemoryDataStore) GetUserCount() int {
	return len(datastore.Users)
}

func (datastore *InMemoryDataStore) GetUserBySeq(seq int) model.User {
	return datastore.Users[seq]
}

func (datastore *InMemoryDataStore) AddUser(user model.User) (int, int) {
	seq := len(datastore.Users) + 1
	user.Seq = seq
	log.Println(user)
	datastore.Users[seq] = user
	return seq, 1
}
