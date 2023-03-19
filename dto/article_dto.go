package dto

import (
	"ginDemo/model"
	"strconv"
)

type ArticleDto struct {
	Id              string `json:"Id"`
	CreatedAt       string `json:"CreatedAt"`
	UpdatedAt       string `json:"UpdatedAt"`
	FileName        string `json:"FileName"`
	UserId          string `json:"UserId"`
	FileStream      string `json:"FileStream"`
	ParentArticleId string `json:"ParentArticleId"`
	Description     string `json:"Description"`
}

func ToArticleDto(article model.Article) ArticleDto {
	return ArticleDto{
		Id:              strconv.FormatInt(int64(article.ID), 10),
		CreatedAt:       article.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:       article.UpdatedAt.Format("2006-01-02 15:04:05"),
		FileName:        article.FileName,
		UserId:          article.UserId,
		FileStream:      article.FileStream,
		ParentArticleId: article.ParentArticleId,
		Description:     article.Description,
	}
}

func ToArticlesDto(articles *[]model.Article) []ArticleDto {
	var articlesDto []ArticleDto
	for _, article := range *articles {
		articleDto := ToArticleDto(article)
		articlesDto = append(articlesDto, articleDto)
	}

	return articlesDto
}
