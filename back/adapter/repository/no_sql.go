package repository

import "devport/domain/repo/nosql"

type NoSQL interface {
	UserRepository() nosql.UserRepository
}
