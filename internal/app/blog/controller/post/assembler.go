package post

import (
	"nothing/internal/app/blog/model/post"
	"nothing/internal/app/blog/model/user"
)

type PostAssembler struct {
}

func NewAssembler() *PostAssembler {
	return &PostAssembler{}
}

func (pa *PostAssembler) AssemblePostBoToVoForHiddenList(boList []*post.PostBo) []*post.HiddenPostVo {
	var voList []*post.HiddenPostVo
	for _, bo := range boList {
		vo := &post.HiddenPostVo{
			ID:        bo.ID,
			CreatedAt: bo.CreatedAt,
			Title:     bo.Title,
		}
		voList = append(voList, vo)
	}
	return voList
}

func (pa *PostAssembler) AssemblePostBoToVoForSimpleList(boList []*post.PostBo) []*post.SimplePostVo {
	var voList []*post.SimplePostVo
	for _, bo := range boList {
		vo := &post.SimplePostVo{
			ID:         bo.ID,
			CreatedAt:  bo.CreatedAt,
			Title:      bo.Title,
			UpdatedAt:  bo.UpdatedAt,
			Type:       bo.Type,
			Summary:    bo.Summary,
			Tags:       bo.Tags,
			Location:   bo.Location,
			ExtendID:   bo.ExtendID,
			CategoryID: bo.CategoryID,
			UserID:     bo.UserID,
			Category: post.PostCategoryVo{
				ID:   bo.Category.ID,
				Name: bo.Category.Name,
			},
			Extend: post.PostExtendVo{
				ID:    bo.Extend.ID,
				Likes: bo.Extend.Likes,
				Views: bo.Extend.Views,
			},
			User: user.UserVo{
				ID:        bo.User.ID,
				Username:  bo.User.Username,
				Role:      bo.User.Role,
				Avatar:    bo.User.Avatar,
				CreatedAt: bo.User.CreatedAt,
			},
		}
		voList = append(voList, vo)
	}
	return voList
}

func (pa *PostAssembler) AssemblePostBoToVoForNormalList(boList []*post.PostBo, respAttachmentMode int) []*post.NormalPostVo {
	var voList []*post.NormalPostVo
	for _, bo := range boList {
		vo := &post.NormalPostVo{
			ID:         bo.ID,
			CreatedAt:  bo.CreatedAt,
			Title:      bo.Title,
			UpdatedAt:  bo.UpdatedAt,
			Type:       bo.Type,
			Summary:    bo.Summary,
			Tags:       bo.Tags,
			Location:   bo.Location,
			ExtendID:   bo.ExtendID,
			CategoryID: bo.CategoryID,
			UserID:     bo.UserID,
			Category: post.PostCategoryVo{
				ID:   bo.Category.ID,
				Name: bo.Category.Name,
			},
			Extend: post.PostExtendVo{
				ID:    bo.Extend.ID,
				Likes: bo.Extend.Likes,
				Views: bo.Extend.Views,
			},
			User: user.UserVo{
				ID:        bo.User.ID,
				Username:  bo.User.Username,
				Role:      bo.User.Role,
				Avatar:    bo.User.Avatar,
				CreatedAt: bo.User.CreatedAt,
			},
			AttachmentMap: func() map[int][]post.PostAttachmentVo {
				if respAttachmentMode == 0 {
					return nil
				}
				return pa.AssembleAttachmentToMap(bo.Attachments)
			}(),
			Attachments: pa.AssembleAttachment(bo.Attachments),
		}
		voList = append(voList, vo)
	}
	return voList
}

func (pa *PostAssembler) AssembleAttachment(boList []post.PostAttachment) []post.PostAttachmentVo {
	var voList []post.PostAttachmentVo
	for _, bo := range boList {
		vo := post.PostAttachmentVo{
			ID:            bo.ID,
			CreatedAt:     bo.CreatedAt,
			PrimaryType:   bo.PrimaryType,
			SecondaryType: bo.SecondaryType,
			Content:       bo.Content,
			PostID:        bo.PostID,
		}
		voList = append(voList, vo)
	}
	return voList
}

func (pa *PostAssembler) AssembleAttachmentToMap(boList []post.PostAttachment) map[int][]post.PostAttachmentVo {
	voMap := make(map[int][]post.PostAttachmentVo)
	for _, bo := range boList {
		//if item, ok := voMap[bo.PrimaryType]; ok {
		vo := post.PostAttachmentVo{
			ID:            bo.ID,
			CreatedAt:     bo.CreatedAt,
			PrimaryType:   bo.PrimaryType,
			SecondaryType: bo.SecondaryType,
			Content:       bo.Content,
			PostID:        bo.PostID,
		}
		voMap[bo.PrimaryType] = append(voMap[bo.PrimaryType], vo)
		//}

	}
	return voMap
}

func (pa *PostAssembler) AssemblePartitionPostBoToVo(boList []*post.PostPartitionBo) []*post.PartitionPostVo {
	var voList []*post.PartitionPostVo
	for _, bo := range boList {
		vo := &post.PartitionPostVo{
			ID:        bo.ID,
			CreatedAt: bo.CreatedAt,
			Title:     bo.Title,
			//UpdatedAt:  bo.UpdatedAt,
			Type:         bo.Type,
			Summary:      bo.Summary,
			Tags:         bo.Tags,
			Location:     bo.Location,
			ExtendID:     bo.ExtendID,
			CategoryID:   bo.CategoryID,
			CategoryName: bo.CategoryName,
			UserID:       bo.UserID,
			//Category: post.PostCategoryVo{
			//	ID:   bo.Category.ID,
			//	Name: bo.Category.Name,
			//},
			//Extend: post.PostExtendVo{
			//	ID:    bo.Extend.ID,
			//	Likes: bo.Extend.Likes,
			//	Views: bo.Extend.Views,
			//},
			//User: user.UserVo{
			//	ID:        bo.User.ID,
			//	Username:  bo.User.Username,
			//	Role:      bo.User.Role,
			//	Avatar:    bo.User.Avatar,
			//	CreatedAt: bo.User.CreatedAt,
			//},
		}
		voList = append(voList, vo)
	}
	return voList
}
