package repository

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

type Topic struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}
type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}
func (*TopicDao) QueryTopicById(id int64) *Topic {
	mu.RLock()
	defer mu.RUnlock()
	return topicIndexMap[id]
}

func (*TopicDao) CreateTopic(title, content string) error {
	mu.Lock()
	defer mu.Unlock()

	newTopic := Topic{Id: curId, Title: title, Content: content, CreateTime: time.Now().Unix()}
	topicIndexMap[curId] = &newTopic
	curId++
	file, err := os.OpenFile(filePath+"post", os.O_APPEND, os.ModeAppend)
	if err != nil {
		return err
	}
	temp, err := json.Marshal(newTopic)
	if err != nil {
		return err
	}
	file.Write([]byte{'\n'})
	file.Write(temp)

	return nil
}
