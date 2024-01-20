package repository

type Database struct {
	Name string
}

type DatabasePostgreSQL Database
type DatabaseMongo Database

func NewPostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{Name: "POSTGRESQL"})
}

func NewMongo() *DatabaseMongo {
	return (*DatabaseMongo)(&Database{Name: "MONGODB"})
}

type DatabaseRepository struct {
	*DatabasePostgreSQL
	*DatabaseMongo
}

func NewDBRepository(sql *DatabasePostgreSQL, mongo *DatabaseMongo) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePostgreSQL: sql,
		DatabaseMongo:      mongo,
	}
}
