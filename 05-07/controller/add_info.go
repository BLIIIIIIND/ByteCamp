package controller

import (
	"github.com/BLIIIIIIND/ByteCamp/05-07/service"
	"strconv"
)

func Create(_parentId, title, content string) error {
	parentId, err := strconv.ParseInt(_parentId, 10, 64)
	if err != nil {
		return err
	}
	if parentId != 0 {
		return service.CreatePost(parentId, content)
	}
	return service.CreateTopic(title, content)
}
