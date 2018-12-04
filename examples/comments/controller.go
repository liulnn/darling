package controllers

import (
	"darling"
	"encoding/json"
	"strconv"
)

type CommentEntity struct {
	Content string
}

type CommentsController struct {
	darling.Controller
	TopicId string
}

func (c *CommentsController) Prepare() {
	c.TopicId = c.PathParams[0]
}

func (c *CommentsController) Post() {
	req := c.Request.In
	req.ParseForm()
	name := req.PostFormValue("name")
	fmt.Println("name:", name)

	c.Response.StatusCode = 201
}

func (c *CommentsController) Get() {
	pageSize, pageNum := 20, 1
	
}

type CommentController struct {
	darling.Controller
	TopicId   string
	CommentId int64
}

func (c *CommentController) Prepare() {
	
}

func (c *CommentController) Get() {
	

}

func (c *CommentController) Delete() {
	
}
