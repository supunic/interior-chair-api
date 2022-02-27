package port

import (
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
	FetchAll() ([]*data.AuthorRepositoryData, error)
}
