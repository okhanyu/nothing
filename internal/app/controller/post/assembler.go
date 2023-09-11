package post

import (
	post2 "nothing/internal/app/model/post"
	"nothing/internal/app/model/user"
)

type PostAssembler struct {
}

func NewAssembler() *PostAssembler {
	return &PostAssembler{}
}

func (pa *PostAssembler) AssemblePostBoToVoForHiddenList(boList []*post2.PostBo) []*post2.HiddenPostVo {
	var voList []*post2.HiddenPostVo
	for _, bo := range boList {
		vo := &post2.HiddenPostVo{
			ID:        bo.ID,
			CreatedAt: bo.CreatedAt,
			Title:     bo.Title,
		}
		voList = append(voList, vo)
	}
	return voList
}

func (pa *PostAssembler) AssemblePostBoToVoForSimpleList(boList []*post2.PostBo) []*post2.SimplePostVo {
	var voList []*post2.SimplePostVo
	for _, bo := range boList {
		vo := &post2.SimplePostVo{
			ID:         bo.ID,
			CreatedAt:  bo.CreatedAt,
			Title:      bo.Title,
			UpdatedAt:  bo.UpdatedAt,
			Type:       bo.Type,
			Summary:    bo.Summary,
			Tags:       bo.Tags,
			Location:   bo.Location,
			CategoryID: bo.CategoryID,
			UserID:     bo.UserID,
			Category: post2.PostCategoryVo{
				ID:   bo.Category.ID,
				Name: bo.Category.Name,
			},
			Extend: post2.PostExtendVo{
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

func (pa *PostAssembler) AssemblePostBoToVoForNormalList(boList []*post2.PostBo, respAttachmentMode int) []*post2.NormalPostVo {
	var voList []*post2.NormalPostVo
	for _, bo := range boList {
		var vo *post2.NormalPostVo
		if bo.Hide == 0 {
			vo = &post2.NormalPostVo{
				ID:         bo.ID,
				CreatedAt:  bo.CreatedAt,
				Title:      bo.Title,
				UpdatedAt:  bo.UpdatedAt,
				Type:       bo.Type,
				Summary:    bo.Summary,
				Tags:       bo.Tags,
				Location:   bo.Location,
				CategoryID: bo.CategoryID,
				UserID:     bo.UserID,
				Hide:       bo.Hide,
				Category: post2.PostCategoryVo{
					ID:   bo.Category.ID,
					Name: bo.Category.Name,
				},
				Extend: post2.PostExtendVo{
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
				AttachmentMap: func() map[int][]post2.PostAttachmentVo {
					if respAttachmentMode == 0 {
						return nil
					}
					return pa.AssembleAttachmentToMap(bo.Attachments)
				}(),
				Attachments: pa.AssembleAttachment(bo.Attachments),
			}
		} else {
			vo = &post2.NormalPostVo{
				ID:        bo.ID,
				CreatedAt: bo.CreatedAt,
				Title:     bo.Title,
				Type:      bo.Type,
				Hide:      bo.Hide,
			}
		}

		voList = append(voList, vo)
	}
	return voList
}

func (pa *PostAssembler) AssembleAttachment(boList []post2.PostAttachment) []post2.PostAttachmentVo {
	var voList []post2.PostAttachmentVo
	for _, bo := range boList {
		vo := post2.PostAttachmentVo{
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

func (pa *PostAssembler) AssembleAttachmentToMap(boList []post2.PostAttachment) map[int][]post2.PostAttachmentVo {
	voMap := make(map[int][]post2.PostAttachmentVo)
	for _, bo := range boList {
		//if item, ok := voMap[bo.PrimaryType]; ok {
		vo := post2.PostAttachmentVo{
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

func (pa *PostAssembler) AssemblePartitionPostBoToVo(boList []*post2.PostPartitionBo) []*post2.PartitionPostVo {
	var voList []*post2.PartitionPostVo
	for _, bo := range boList {
		vo := &post2.PartitionPostVo{
			ID:        bo.ID,
			CreatedAt: bo.CreatedAt,
			Title:     bo.Title,
			//UpdatedAt:  bo.UpdatedAt,
			Type:         bo.Type,
			Summary:      bo.Summary,
			Tags:         bo.Tags,
			Location:     bo.Location,
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
