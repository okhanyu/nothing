package post

import (
	"golang.org/x/net/context"
	"log"
	"nothing/config"
	pkgPost "nothing/internal/app/model/post"
	pkgPostRepo "nothing/internal/app/repository/post"
	"nothing/pkg/util"
)

const (
	entityAttachment = "attachments"
	entityCategory   = "category"
	entityUser       = "user"
	entityExtend     = "extend"
)

type PostService struct {
	Repository pkgPostRepo.PostRepository
}

func NewPostService(rep pkgPostRepo.PostRepository) *PostService {
	return &PostService{
		Repository: rep,
	}
}

func (ps *PostService) FindByID(id int64) (*pkgPost.PostBo, error) {
	postInfo, err := ps.Repository.FindByID(
		id,
		pkgPostRepo.WithAttachment(),
		pkgPostRepo.WithCategory(),
		pkgPostRepo.WithUser(),
		pkgPostRepo.WithExtend(),
	)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return postInfo, nil
}

func (ps *PostService) FindBatchPartition(req pkgPost.FindReq) ([]*pkgPost.PostPartitionBo, error) {
	partition, err := ps.Repository.FindBatchPartition(req, config.Global.Business.RowNum)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return partition, nil
}

func (ps *PostService) FindBatch(req pkgPost.FindReq) ([]*pkgPost.PostBo, error) {
	rep := ps.Repository

	selects, entities := pkgPostRepo.FieldsCheck(req.Fields)

	var posts []*pkgPost.PostBo
	var err error

	posts, err = rep.FindBatch(
		func() pkgPostRepo.Option {
			if ok := util.ArraysStringContain(entities, entityAttachment); ok {
				return pkgPostRepo.WithAttachment()
			}
			return nil
		}(),
		func() pkgPostRepo.Option {
			if ok := util.ArraysStringContain(entities, entityCategory); ok {
				return pkgPostRepo.WithCategory()
			}
			return nil
		}(),
		func() pkgPostRepo.Option {
			if ok := util.ArraysStringContain(entities, entityUser); ok {
				return pkgPostRepo.WithUser()
			}
			return nil
		}(),
		func() pkgPostRepo.Option {
			if ok := util.ArraysStringContain(entities, entityExtend); ok {
				return pkgPostRepo.WithExtend()
			}
			return nil
		}(),
		pkgPostRepo.Select(selects),
		pkgPostRepo.WhereCategory(req.Filters.CategoryId),
		pkgPostRepo.WhereType(req.Filters.Type),
		pkgPostRepo.WhereHide(req.Filters.Hide),
		pkgPostRepo.Order(pkgPostRepo.PostSort, req.Order),
		pkgPostRepo.Limit(req.Page.PageNumber, req.Page.PageSize),
	)

	//switch req.Mode {
	//case pkgPostCons.Normal:
	//	posts, err = rep.FindBatch(
	//		pkgPostRepo.WithAttachment(),
	//		pkgPostRepo.WithCategory(),
	//		pkgPostRepo.WithUser(),
	//		pkgPostRepo.WithExtend(),
	//		pkgPostRepo.WhereCategory(req.CategoryId),
	//		pkgPostRepo.WhereType(req.Type),
	//		pkgPostRepo.WhereHide(nil),
	//		pkgPostRepo.Order(pkgPostRepo.PostSort, req.Order),
	//		pkgPostRepo.Limit(req.Page.PageNumber, req.Page.PageSize),
	//	)
	//	break
	//case pkgPostCons.Simple:
	//	posts, err = rep.FindBatch(
	//		//repo.Select([]string{"id", "user_id", "title", "category_id", "tags", "location", "created_at"}),
	//		pkgPostRepo.WithCategory(),
	//		pkgPostRepo.WithUser(),
	//		pkgPostRepo.WithExtend(),
	//		pkgPostRepo.WhereCategory(req.CategoryId),
	//		pkgPostRepo.WhereType(req.Type),
	//		pkgPostRepo.WhereHide(nil),
	//		pkgPostRepo.Order(pkgPostRepo.PostSort, req.Order),
	//		pkgPostRepo.Limit(req.Page.PageNumber, req.Page.PageSize),
	//	)
	//	break
	//case pkgPostCons.Hidden:
	//	posts, err = rep.FindBatch(
	//		pkgPostRepo.Select([]string{"id", "title", "created_at"}),
	//		pkgPostRepo.WithCategory(),
	//		pkgPostRepo.WithUser(),
	//		pkgPostRepo.WithExtend(),
	//		pkgPostRepo.WhereCategory(req.CategoryId),
	//		pkgPostRepo.WhereType(req.Type),
	//		pkgPostRepo.WhereHide([]int{1}),
	//		pkgPostRepo.Order(pkgPostRepo.PostSort, req.Order),
	//		pkgPostRepo.Limit(req.Page.PageNumber, req.Page.PageSize),
	//	)
	//	break
	//}

	if err != nil {
		log.Print(err)
		return nil, err
	}
	return posts, nil
}

func (ps *PostService) CreatePost(ctx context.Context, req pkgPost.CreateReq) error {
	err := ps.Repository.CreatePost(ctx, req)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (ps *PostService) FindAttachmentContent(req pkgPost.FindReq) ([]string, error) {
	rep := ps.Repository

	var contents []string
	var err error
	contents, err = rep.FindAttachmentContent(
		req,
		//req, pkgPostRepo.Limit(req.Page.PageNumber, req.Page.PageSize),
	)

	if err != nil {
		log.Print(err)
		return nil, err
	}
	return contents, nil
}
