package repository

import "devport/domain/repository/nosql"

type NoSQL interface {
	UserRepository() nosql.UserRepository
}
