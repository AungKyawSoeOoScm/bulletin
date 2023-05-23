package interfaces

import "gin_test/bulletin_board/model"

type PostsInterface interface {
	Save(tags model.Posts) error
	Update(tags model.Posts)
	Delete(tagsId int)
	FindById(tagsId int) (tags model.Posts, err error)
	FindAll() []model.Posts
}
