package main

import (
	"github.com/BLIIIIIIND/ByteCamp/05-07/controller"
	"github.com/BLIIIIIIND/ByteCamp/05-07/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := Init(); err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := controller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})

	r.POST("community/post", func(c *gin.Context) {
		parentId, title, content := c.PostForm("parent_id"), c.PostForm("title"), c.PostForm("content")
		err := controller.Create(parentId, title, content)
		if err != nil {
			c.JSON(400, err.Error())
		} else {
			c.JSON(200, "Added")
		}
	})

	err := r.Run()
	if err != nil {
		return
	}
}

func Init() error {
	if err := repository.Init(); err != nil {
		return err
	}
	return nil
}
