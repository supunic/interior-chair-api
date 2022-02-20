package port

import (
	"app/entity/model/author"
	"app/usecase/data"
)

type AuthorInputPort interface {
	FetchAll()
}

type AuthorOutputPort interface {
	FetchAll(authorOutputDataList []*data.AuthorOutputData)
	Error(err error)
}

type AuthorRepository interface {
	FetchAll() ([]*author.Author, error)
}
