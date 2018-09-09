# Hacker News
## What is it
Go wrapper for the Hacker News [REST api](https://github.com/HackerNews/API).
The API is based at firebaseio.com

## Installation
>go get github.com/nexes/hackernews

# HN Objects
## Story
The Story struct is an object that will contain information about a single story.

|Field|Description|
|-----|-----------|
|ID|The unique ID Hacker News gives every item|
|Author| The username of who posted the story|
|Title| The title of the story|
|URL| The URL to where the story is found|
|Score| The story's score, the votes given|
|Time| The time the story was created, Unix Time|
|CommentIDs| List of the unique comment IDs for the story|
|CommentCount| The number of comments _including replies to comments_|
|IDType| A string showing if this is a story, if so the value will be "story"|

## ChangedStories
The ChangedStories object will hold a list of IDs to stories and profiles that have changed

|Field|Description|
|-----|-----------|
|ItemIDS| A list of IDs representing changed stories|
|Profiles| A list of IDs representing changed user profiles|

## User
The User object will hold public information about a user

|Field|Description|
|-----|-----------|
|ID| A string showing the unique username|
|Created| The time the users account was created. Unix Time|
|Karma| An integer showing the users current karma|
|About| The users description _not all users will have done this_|
|ActivityIDs| A list of unique IDs for all the users posted stories, comment etc|

## Comment
The Comment object will hold information about a comment to a story

|Field|Description|
|-----|-----------|
|ID| The unique ID for the comment|
|ParentID| The unique ID for the parent. This could be the story or the comment this is replying to|
|Author| The username who submitted the comment|
|Text| The comment body|
|Time| The time the comment was posted. Unix Time|
|Replies| A list of unique IDs the this comments replies. These are IDs to other Comment objects|
|IDType| A string showing if this is a comment, if so the value will be "comment"|

## Constants
* **TopStoryID**
* **NewStoryID**
* **AskStoryID**

The function _GetStoryIDList_ takes an integer parameter. Pass in one of these constants to return the IDs of that type of item.

## Example
```Go
package main

import (
	"fmt"

	"github.com/nexes/hackernews"
)

func main() {
    //ids will be an int slice.Usually the list returned from 
    //GetStoryIDList will be around 500.
    ids, err := hn.GetStoryIDList(hn.TopStoryID)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    //GetStoryFromID makes a GET request and will return a Story
    //object from the json returned. Use goroutines when looping through all
    //ids returned form GetStoryIDList
    story, err := hn.GetStoryFromID(ids[0])
    if err != nil {
        fmt.Println(err)
        return
    }
    
    fmt.Println("First Story's Title = ", story.Title)
    fmt.Println("First Story's Author = ", story.Author)
    fmt.Println("First Story's Karma = ", story.Score)
    
    comments, err := story.GetComments()
    if err != nil {
        fmt.Println(err)
        return
    }
    
    fmt.Println("First Comment's Author = ", comments[0].Author)
    fmt.Println("First Comment's comment = ", comments[0].Text)
    
    user, err := story.GetUser()
    if err != nil {
        fmt.Println(err)
        return
    }
    
    fmt.Println("User's username = ", user.ID)
    fmt.Println("User's karma = ", user.Karma)
    fmt.Println("User's description = ", user.About)
}
```

# LICENSE (MIT)
Copyright (c) 2016 Joe Berria
