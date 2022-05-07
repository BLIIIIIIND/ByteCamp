package service

import "github.com/Moonlight-Zhao/go-project-example/repository"

func CreatePost(parentId int64, content string) error {
	return repository.NewPostDaoInstance().CreatePost(parentId, content)
}

func CreateTopic(title, content string) error {
	return repository.NewTopicDaoInstance().CreateTopic(title, content)
}
