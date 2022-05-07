package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type Post struct {
	Id         int64  `json:"id"`
	ParentId   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}
type PostDao struct {
}

var (
	postDao  *PostDao
	postOnce sync.Once
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}
func (*PostDao) QueryPostsByParentId(parentId int64) []*Post {
	mu.RLock()
	defer mu.RUnlock()
	return postIndexMap[parentId]
}

func (*PostDao) CreatePost(parentId int64, content string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := postIndexMap[parentId]; !ok {
		return fmt.Errorf("no sunch parent topic")
	}
	newPost := Post{Id: curId, ParentId: parentId, Content: content, CreateTime: time.Now().Unix()}
	postIndexMap[parentId] = append(postIndexMap[parentId], &newPost)
	curId++
	file, err := os.OpenFile(filePath+"post", os.O_APPEND, os.ModeAppend)
	if err != nil {
		return err
	}
	temp, err := json.Marshal(newPost)
	if err != nil {
		return err
	}
	file.Write([]byte{'\n'})
	file.Write(temp)

	return nil
}
