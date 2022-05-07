package service

import (
	"github.com/BLIIIIIIND/ByteCamp/05-07/repository"
)

func CreatePost(parentId int64, content string) error {
	return repository.NewPostDaoInstance().CreatePost(parentId, content)
}

func CreateTopic(title, content string) error {
	return repository.NewTopicDaoInstance().CreateTopic(title, content)
}
