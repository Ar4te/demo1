package controller

import (
	"ginDemo/common"
	"ginDemo/model"
	"ginDemo/response"
	"github.com/gin-gonic/gin"
	"time"
)

func PubTopic(c *gin.Context) {
	DB := common.GetDB()
	var topic model.Topic

	if err := c.ShouldBind(&topic); err != nil {
		response.Fail(c, gin.H{"code": 400, "msg": "传值错误"}, "传值错误")
		return
	}

	topic.ID, _ = common.Generate()
	topic.CreatedAt = time.Now()

	res := DB.Create(topic)
	if res.Error != nil || res.RowsAffected <= 0 {
		response.Fail(c, gin.H{"code": 500, "msg": res.Error}, "记录话题失败")
		return
	}

	response.Success(c, gin.H{"code": 200}, "记录话题成功")
}
