package controller

import (
	"github.com/gin-gonic/gin"
	"ginDemo/response"
	"ginDemo/common"
	"ginDemo/model"
	"io/ioutil"
	"encoding/base64"
	"github.com/spf13/viper"
)

func UploadFile (c *gin.Context) {
	DB := common.GetDB()
	name := c.PostForm("name")
	userId := c.PostForm("userId")
	fileName := c.PostForm("fileName")
	file, err := c.FormFile("file")
	if err != nil {
		msg := "get form err: " + err.Error()
		response.Fail(c, gin.H{}, msg)
		// c.JSON(http.StatusBadRequest, gin.H{"code":400,"msg":msg})
		return
	}

	// 将*multipart.FileHeader转成*os.File
	_file, _ := file.Open()
	defer _file.Close()

	// 读取文件存为[]uint8
	fileData, _ := ioutil.ReadAll(_file)

	// 将[]uint8转成base64
	bs64 := base64.StdEncoding.EncodeToString([]byte(fileData))

	var user model.User
	DB.First(&user, userId)

	if user.ID == 0 {
		response.Fail(c, gin.H{}, "用户不存在")
		// c.JSON(http.StatusUnauthorized, gin.H{"code":401,"msg":"权限不足"})
		return
	}

	newArticle := model.Article{
		Name: name,
		UserId: userId,
		FileName: fileName,
		FileStream: bs64,
	}

	DB.Create(&newArticle)

	response.Success(c, gin.H{"name":name, "userId":userId, "fileName": fileName}, "上传成功")
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

	savepath := viper.GetString("filesavepath.path") + "/" + article.FileName + ".txt"
	err = ioutil.WriteFile(savepath, bytes, 0666)

	if err != nil {
		msg := "write file fail, err:" + err.Error()
		response.Fail(c, gin.H{}, msg)
	}

	response.Success(c, gin.H{}, "下载成功")
}