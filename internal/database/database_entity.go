package database

type DatabaseEntity interface {
	GetCollectionName() string
	New() DatabaseEntity
	GetFilterOne() map[string]interface{}
}
