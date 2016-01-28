package hn

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//User is a data structure representing a user
type User struct {
	ID          string `json:"id"`
	Created     int    `json:"created"`
	Karma       int    `json:"karma"`
	About       string `json:"about"`
	ActivityIDs []int  `json:"submitted"`
}

func getUserByID(userName string) (User, error) {
	var user User
	url := fmt.Sprintf("%suser/%s.json", baseURL, userName)

	res, err := http.Get(url)
	if err != nil {
		return user, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return user, err
	}

	json.Unmarshal(data, &user)
	return user, nil
}

//GetUser return a User object for the username passed.
func GetUser(name string) (User, error) {
	return getUserByID(name)
}

//GetStorySubmissions returns an array of story objects of the stories that the user has submitted
func (user *User) GetStorySubmissions() ([]Story, error) {
	var count = len(user.ActivityIDs)
	var stories []Story

	if count < 1 {
		return nil, errors.New("User has no public submissions")
	}

	for i := 0; i < count; i++ {
		story, err := GetStoryFromID(user.ActivityIDs[i])
		if err == nil {
			stories = append(stories, story)
		}
	}

	return stories, nil
}

//GetComments will return the comments left by the user
func (user *User) GetComments() ([]Comment, error) {
	var count = len(user.ActivityIDs)
	var comments []Comment

	if count < 1 {
		return nil, errors.New("User has no public comments")
	}

	for i := 0; i < count; i++ {
		com, err := createCommentFromID(user.ActivityIDs[i])
		if err == nil {
			comments = append(comments, com)
		}
	}

	return comments, nil
}
