package post

import (
	"golang.org/x/net/context"
	"log"
	"nothing/config/blog"
	postcons "nothing/internal/app/blog/cons/post"
	"nothing/internal/app/blog/model/post"
	postRepo "nothing/internal/app/blog/repository/post"
)

type PostService struct {
	Repository postRepo.PostRepository
}

func NewPostService(rep postRepo.PostRepository) *PostService {
	return &PostService{
		Repository: rep,
	}
}

func (ps *PostService) FindByID(id int64) (*post.PostBo, error) {
	postInfo, err := ps.Repository.FindByID(
		id,
		postRepo.WithAttachment(),
		postRepo.WithCategory(),
		postRepo.WithUser(),
		postRepo.WithExtend(),
	)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return postInfo, nil
}

func (ps *PostService) FindBatchPartition(req post.FindReq) ([]*post.PostPartitionBo, error) {
	partition, err := ps.Repository.FindBatchPartition(req, blog.Global.Business.RowNum)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return partition, nil
}

func (ps *PostService) FindBatch(req post.FindReq) ([]*post.PostBo, error) {
	rep := ps.Repository

	var posts []*post.PostBo
	var err error
	switch req.Mode {
	case postcons.Normal:
		posts, err = rep.FindBatch(
			postRepo.WithAttachment(),
			postRepo.WithCategory(),
			postRepo.WithUser(),
			postRepo.WithExtend(),
			postRepo.WhereCategory(req.CategoryId),
			postRepo.WhereType(req.Type),
			postRepo.WhereHide(nil),
			postRepo.Order(postRepo.PostSort, req.Order),
			postRepo.Limit(req.Page.PageNumber, req.Page.PageSize),
		)
		break
	case postcons.Simple:
		posts, err = rep.FindBatch(
			//repo.Select([]string{"id", "user_id", "title", "category_id", "tags", "location", "created_at"}),
			postRepo.WithCategory(),
			postRepo.WithUser(),
			postRepo.WithExtend(),
			postRepo.WhereCategory(req.CategoryId),
			postRepo.WhereType(req.Type),
			postRepo.WhereHide(nil),
			postRepo.Order(postRepo.PostSort, req.Order),
			postRepo.Limit(req.Page.PageNumber, req.Page.PageSize),
		)
		break
	case postcons.Hidden:
		posts, err = rep.FindBatch(
			postRepo.Select([]string{"id", "title", "created_at"}),
			postRepo.WithCategory(),
			postRepo.WithUser(),
			postRepo.WithExtend(),
			postRepo.WhereCategory(req.CategoryId),
			postRepo.WhereType(req.Type),
			postRepo.WhereHide([]int{1}),
			postRepo.Order(postRepo.PostSort, req.Order),
			postRepo.Limit(req.Page.PageNumber, req.Page.PageSize),
		)
		break
	}

	if err != nil {
		log.Print(err)
		return nil, err
	}
	return posts, nil
}

func (ps *PostService) CreatePost(ctx context.Context, req post.CreateReq) error {
	err := ps.Repository.CreatePost(ctx, req)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (ps *PostService) FindAttachmentContent(req post.FindReq) ([]string, error) {
	rep := ps.Repository

	var contents []string
	var err error
	contents, err = rep.FindAttachmentContent(
		req,
		//req, postRepo.Limit(req.Page.PageNumber, req.Page.PageSize),
	)

	if err != nil {
		log.Print(err)
		return nil, err
	}
	return contents, nil
}
