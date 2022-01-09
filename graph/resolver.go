package graph

import "sync"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type DB = map[string][]string

type Resolver struct {
	sync.RWMutex
	Db DB
}
