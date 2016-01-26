package hn

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Comment is a data structure representing a comment to a story
type Comment struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parent"`
	Author   string `json:"by"`
	Text     string `json:"text"`
	Time     int    `json:"time"`
	Replies  []int  `json:"kids"`
	IDtype   string `json:"type"`
}

func createCommentFromID(id int) (Comment, error) {
	var comment Comment
	url := fmt.Sprintf("%sitem/%d.json", baseURL, id)

	res, err := http.Get(url)
	if err != nil {
		return comment, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return comment, err
	}

	json.Unmarshal(data, &comment)

	if comment.IDtype == "comment" {
		return comment, nil
	}

	return Comment{}, errors.New("ID that was given is not a comment")
}

//GetUser will return a User object of the comment.
func (comment *Comment) GetUser() (User, error) {
	return getUserByID(comment.Author)
}
