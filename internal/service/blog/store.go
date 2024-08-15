package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/internal/db"
	"github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/internal/model"
)

type BlogStore struct {
	db *db.Rdb
}

func NewBlogStore(db *db.Rdb) *BlogStore {
	return &BlogStore{
		db: db,
	}
}

func (store *BlogStore) FindAll(ctx *gin.Context) (*[]model.Blog, error) {
	var blogs []model.Blog
	tx := store.db.Db.Where("deleted_at is NULL").Scopes(model.Paginate(ctx)).Order("updated_at desc").Find(&blogs)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &blogs, nil
}

func (store *BlogStore) Find(ctx *gin.Context, id uint64) (*model.Blog, error) {
	var blog model.Blog
	tx := store.db.Db.Where("id = ?", id).First(&blog)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &blog, nil
}
