package datastore

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/cream-lab/vanilla/model"
)

type DataStoreFactory func(conf map[string]interface{}) (DataStore, error)

var inMemoryDataStore = &InMemoryDataStore{
	Users: make(map[int]model.User),
}

func NewInMemoryDataStore(conf map[string]interface{}) (DataStore, error) {
	return inMemoryDataStore, nil
}

var datastoreFactories = make(map[string]DataStoreFactory)

func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Panicf("DataStore factory %s does not exist.", name)
	}

	_, registered := datastoreFactories[name]
	if registered {
		log.Panicf("DataStore factory %s already registered. Ignoring", name)
	}

	datastoreFactories[name] = factory
}

func CreateDataStore(conf map[string]interface{}) (DataStore, error) {
	engineName, _ := conf["DATASTORE"].(string)

	engineFactory, ok := datastoreFactories[engineName]
	if !ok {
		availableDatastores := make([]string, len(datastoreFactories))
		for k, _ := range datastoreFactories {
			availableDatastores = append(availableDatastores, k)
		}

		return nil, errors.New(fmt.Sprintf("Invalid Datastore name. Must be one of: %s", strings.Join(availableDatastores, ", ")))
	}

	return engineFactory(conf)
}
