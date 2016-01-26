package hn

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseURL = "https://hacker-news.firebaseio.com/v0/"
)

//Story is the data structure to hold information on a single story=
type Story struct {
	ID           int    `json:"id"`
	Author       string `json:"by"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	Score        int    `json:"score"`
	Time         int    `json:"time"`
	CommentIDs   []int  `json:"kids"`
	CommentCount int    `json:"descendants"`
	IDtype       string `json:"type"`
}

//ChangedStories is a data structure to hold stories and profiles that have changed
type ChangedStories struct {
	ItemIDs  []int    `json:"items"`
	Profiles []string `json:"profiles"`
}

//GetTopStoryIDs will return a slice of IDs Hacker News uses to represent a story.
//this can be refactored to use a common function
func GetTopStoryIDs() ([]int, error) {
	var ids []int
	url := baseURL + "topstories.json"

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &ids)
	return ids, nil
}

//GetNewStoryIDs will return a slice of IDs Hacker News uses to represent a new story.
//this can be refactored to use a common function
func GetNewStoryIDs() ([]int, error) {
	var ids []int
	url := baseURL + "newstories.json"

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &ids)
	return ids, nil
}

//GetUpdatedStories will return an object that has a list of changed stories and profiles
//this can be refactored to use a common function
func GetUpdatedStories() (ChangedStories, error) {
    var changes ChangedStories
    var url = fmt.Sprintf("%supdates.json", baseURL)
    
    res, err := http.Get(url)
    if err != nil {
        return changes, err
    }
    defer res.Body.Close()
    
    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return changes, err
    }
    
    json.Unmarshal(data, &changes)
    return changes, nil
}


// GetStoryFromID will return an Story object from the hacker news ID that is passed
// This function will return an empty Story object if the ID passed was not an ID to a story
func GetStoryFromID(id int) (Story, error) {
	var story Story
	url := fmt.Sprintf("%sitem/%d.json", baseURL, id)

	res, err := http.Get(url)
	if err != nil {
		return story, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return story, err
	}

	json.Unmarshal(data, &story)

	fmt.Println("item id = ", story.IDtype)
	if story.IDtype == "story" {
		return story, nil
	}

	return Story{}, errors.New("ID passed was not an ID to a story")
}

//GetComments return a list of Comment objects representing the stories comments
func (story *Story) GetComments() ([]Comment, error) {
	var comments []Comment
	var length = len(story.CommentIDs)

	for i := 0; i < length; i++ {
		comment, err := createCommentFromID(story.CommentIDs[i])
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
		var replyCount = len(comment.Replies)

		for j := 0; j < replyCount; j++ {
			reply, err := createCommentFromID(comment.Replies[j])
			if err != nil {
				return nil, err
			}
			comments = append(comments, reply)
		}
	}

	return comments, nil
}

//GetUser will return a User object of the comment.
func (story *Story) GetUser() (User, error) {
	return getUserByID(story.Author)
}
