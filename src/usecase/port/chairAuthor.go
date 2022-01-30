package port

import (
	"app/entity/model/chairAuthor"
	"app/usecase/data"
)

type ChairAuthorInputPort interface {
	FetchAll()
}

type ChairAuthorOutputPort interface {
	FetchAll(chairAuthorOutputDataList []*data.ChairAuthorOutputData)
	Error(err error)
}

type ChairAuthorRepository interface {
	FetchAll() ([]*chairAuthor.ChairAuthor, error)
}
