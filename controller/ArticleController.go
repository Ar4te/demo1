package controller

import (
	"encoding/base64"
	"ginDemo/common"
	"ginDemo/model"
	"ginDemo/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"time"
)

func UploadFile(c *gin.Context) {
	DB := common.GetDB()
	userId := c.PostForm("userId")
	file, header, err := c.Request.FormFile("file")
	// 1、multipart.File 是文件对象
	// 2、multipart.FileHeader文件头部包含了一些基本信息
	/*
		type FileHeader struct {
			Filename string                 //文件全称，带扩展名
			Header   textproto.MIMEHeader   //MIME信息
			Size     int64                  //文件大小,单位bit
			content []byte                  //文件内容,类型[]byte
			tmpfile string                  //临时文件
		}
	*/
	if err != nil {
		msg := "get form err: " + err.Error()
		response.Fail(c, gin.H{}, msg)
		return
	}

	buffer := make([]byte, header.Size)
	// 丢弃读取字节数
	_, err = file.Read(buffer)

	if err != nil {
		msg := "read file fail, err: " + err.Error()
		response.Fail(c, gin.H{}, msg)
		return
	}
	// 将[]byte转成base64
	bs64 := base64.StdEncoding.EncodeToString(buffer)

	var user model.User
	DB.First(&user, userId)

	if user.ID == 0 {
		response.Fail(c, gin.H{}, "用户不存在")
		return
	}

	id, _ := common.Generate()

	newArticle := model.Article{
		ID:         id,
		CreatedAt:  time.Now(),
		UserId:     userId,
		FileName:   header.Filename,
		FileStream: bs64,
	}

	DB.Create(&newArticle)

	response.Success(c, gin.H{"userId": userId, "fileName": header.Filename}, "上传成功")
}

func DownloadFile(c *gin.Context) {
	DB := common.GetDB()
	fileid := c.Query("fileid")
	filename := c.Query("filename")

	var article model.Article
	DB.Where("id = ?", fileid).First(&article)
	if article.ID == 0 {
		DB.Where("file_name = ?", filename).First(&article)
		if article.ID == 0 {
			response.Fail(c, gin.H{}, "文件不存在")
			return
		}
	}

	bytes, err := base64.StdEncoding.DecodeString(article.FileStream)
	if err != nil {
		msg := "decode file stream fail, err:" + err.Error()
		response.Fail(c, gin.H{}, msg)
	}

	savepath := viper.GetString("filesavepath.path") + "/" + article.FileName
	err = ioutil.WriteFile(savepath, bytes, 0666)

	if err != nil {
		msg := "write file fail, err:" + err.Error()
		response.Fail(c, gin.H{}, msg)
	}

	response.Success(c, gin.H{"filename": article.FileName, "filestream": article.FileStream}, "下载成功")
}

func GetAllArticle(c *gin.Context) {
	DB := common.GetDB()

	var articles []model.Article
	res := DB.Table("articles").Select([]string{"id", "file_name", "parent_article_id", "created_at", "user_id"}).Find(&articles)

	if res.Error != nil {
		response.Fail(c, gin.H{}, res.Error.Error())
	}
	response.Success(c, gin.H{"dataCount": res.RowsAffected, "data": res.Value}, "true")
}
