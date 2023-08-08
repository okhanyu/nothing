package post

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"nothing/internal/app/blog/model/post"
	repository2 "nothing/internal/app/blog/repository"
	"nothing/internal/pkg/database"
	"strings"
)

const (
	Default         = 0
	AttachmentsSort = "id asc"
	PostSort        = "created_at"
)

type PostRepository interface {
	FindBatch(option ...Option) ([]*post.PostBo, error)
	FindByID(id int64, option ...Option) (*post.PostBo, error)
	FindBatchPartition(req post.FindReq, rowNum int) ([]*post.PostPartitionBo, error)
}

type PostRepositoryImpl struct {
	DB *database.DataBase
}

func NewPostRepository(db *database.DataBase) PostRepository {
	return &PostRepositoryImpl{db}
}

type Option func(db *gorm.DB) *gorm.DB

func WhereCategory(categoryID []int) Option {
	return func(db *gorm.DB) *gorm.DB {
		if categoryID != nil && len(categoryID) != 0 {
			db = db.Where("category_id in ?", categoryID)
		}
		return db
	}
}

func WhereType(postType []int) Option {
	return func(db *gorm.DB) *gorm.DB {
		if postType != nil && len(postType) != 0 {
			db = db.Where("type in ?", postType)
		}
		return db
	}
}

func WhereHide(hide []int) Option {
	return func(db *gorm.DB) *gorm.DB {
		if hide != nil && len(hide) != 0 {
			db = db.Where("hide in ?", hide)
		} else {
			db = db.Where("hide = ?", 0)
		}
		return db
	}
}

func Order(column string, sort int) Option {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Order(func() string {
			if column == "" {
				column = "created_at"
			}
			col := fmt.Sprintf("%s.%s", repository2.PostTable, column)
			if sort == 0 {
				return fmt.Sprintf("%s %s", col, "desc")
			}
			return fmt.Sprintf("%s %s", col, "asc")
		}())
		return db
	}
}

func Limit(pageNumber int, pageSize int) Option {
	return func(db *gorm.DB) *gorm.DB {
		if pageSize <= 0 {
			pageSize = 5
		}
		if pageSize > 100 {
			pageSize = 100
		}
		db = db.Offset(func() int {
			if pageNumber <= 0 {
				return 0
			}
			return pageNumber - 1
		}() * pageSize).Limit(pageSize)
		return db
	}
}

func Select(columns []string) Option {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Select(func() string {
			if columns == nil || len(columns) == 0 {
				return "*"
			}
			return strings.TrimRight(strings.Join(columns, ","), ",")
		}())
		return db
	}
}

func WithAttachment() Option {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Preload(post.Attachments, func(db *gorm.DB) *gorm.DB {
			return db.Order(AttachmentsSort)
		})
		return db
	}
}
func WithUser() Option {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Preload(post.User)
		return db
	}
}
func WithExtend() Option {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Preload(post.Extend)
		return db
	}
}

func WithCategory() Option {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Preload(post.Category)
		return db
	}
}

func (pri *PostRepositoryImpl) FindByID(id int64, opts ...Option) (*post.PostBo, error) {
	db := pri.DB.GormDB
	for _, opt := range opts {
		db = opt(db)
	}
	var postDetail *post.PostBo
	err := db.Where("id = ?", id).Where("deleted = 0").Where("hide = 0").Find(&postDetail).Error
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return postDetail, nil
}

func (pri *PostRepositoryImpl) FindBatch(opts ...Option) ([]*post.PostBo, error) {

	db := pri.DB.GormDB
	for _, opt := range opts {
		db = opt(db)
	}
	var postBoList []*post.PostBo
	err := db.Where("deleted = 0").Find(&postBoList).Error
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return postBoList, nil
}

func (pri *PostRepositoryImpl) FindBatchPartition(req post.FindReq, rowNum int) ([]*post.PostPartitionBo,
	error) {
	var postList []*post.PostPartitionBo
	//db := NormalConditionHandle(s.db, req, QueryBatch)
	err := pri.DB.GormDB.Raw("SELECT subquery.*,pc.name as `category_name` FROM ( SELECT *, "+
		"ROW_NUMBER() OVER ( PARTITION BY category_id ORDER BY id ) AS row_num FROM "+repository2.PostTable+" WHERE type in ? AND deleted = 0 And hide = 0 ORDER BY created_at DESC) "+
		"AS subquery LEFT JOIN "+repository2.PostCategoryTable+" AS pc ON pc.`id` = subquery."+
		"category_id WHERE row_num <= ? ",
		req.Type, rowNum).Scan(&postList).Error
	if err != nil {
		panic(err)
	}

	fmt.Println(postList)

	if err != nil {
		log.Print(err)
		return nil, err
	}
	return postList, nil
}
