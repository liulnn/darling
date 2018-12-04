package comments

import (
	"time"
)

type App struct {
	Id          int64
	Secret      string
	Name        string
	PackageName string
}

func AddApp(name, packageName string) (id int64, err error) {
	
}

func ValidateApp(id int64, secret string) bool {

}

type Comment struct {
	Id        int64 
	Topic     string    
	Source    string    
	Username  string   
	Content   string   
	CreatedAt time.Time 
}


func AddComment(topicId, source, username, content string) (id int64, err error) {
	
}

func GetComments(topicId string, pageSize int, pageNum int) (comments []*Comment, err error) {

}

func GetComment(topicId string, id int64) (c *Comment, err error) {

}

func DeleteComment(topicId string, id int64) (success bool, err error) {

}
