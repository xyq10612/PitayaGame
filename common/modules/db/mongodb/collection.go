package mongodb

type Collection interface {
	GetDBName() string
	GetCollectionName() string
}
