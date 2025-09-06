/*
 * Copyright (c) 2025 Michael Plunkett (https://github.com/michplunkett)
 * All rights reserved.
 * Used to delete Reddit user posts and comments older than a given threshold date.
 */

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type envVars struct {
	appID        string
	appSecret    string
	userName     string
	userPassword string
}

const LIMIT = 100

var THRESHOLD = time.Now().AddDate(-2, 0, 0)

func arrayHasNoEmptyStrings(envVars []string) bool {
	for _, value := range envVars {
		if value == "" {
			return false
		}
	}

	return true
}

func main() {
	ctx := context.Background()

	e := envVars{
		os.Getenv("REDDIT_APP_ID"),
		os.Getenv("REDDIT_SECRET"),
		os.Getenv("REDDIT_USER_ID"),
		os.Getenv("REDDIT_USER_PASSWORD"),
	}

	if !arrayHasNoEmptyStrings([]string{e.appID, e.appSecret, e.userName, e.userPassword}) {
		log.Fatal(fmt.Errorf("one of the last.fm environment variables is not present in your system"))
	}

	credentials := reddit.Credentials{ID: e.appID, Secret: e.appSecret, Username: e.userName, Password: e.userPassword}
	cli, clientErr := reddit.NewClient(credentials)
	if clientErr != nil {
		log.Fatal(clientErr)
	}

	// Get overview of user
	_ = *cli.Comment
	_ = *cli.Post
	userService := *cli.User

	// Get all posts
	fmt.Println("------- Pulling user posts")

	lastPostID := ""
	postIds := make([]string, 0)
	postOptions := &reddit.ListUserOverviewOptions{
		ListOptions: reddit.ListOptions{
			Limit: LIMIT,
		},
		Sort: "new",
		Time: "all",
	}
	for {
		if lastPostID != "" {
			postOptions.ListOptions.After = lastPostID
		}

		posts, _, err := userService.Posts(ctx, postOptions)
		if err != nil {
			log.Fatal(err)
		}

		if len(posts) == 0 {
			break
		}

		for _, post := range posts {
			if post.Created.Time.Before(THRESHOLD) ||
				post.Created.Time.Equal(THRESHOLD) {
				postIds = append(postIds, post.FullID)
			}
		}

		lastPostID = posts[len(posts)-1].FullID
	}

	// Delete all posts
	fmt.Println("------- Deleting user posts:", len(postIds))
	for _, pID := range postIds {
		fmt.Println(pID)
		//_, err := postService.Delete(ctx, pID)
		//if err != nil {
		//	log.Fatal(err)
		//}
	}

	// Get all comments
	fmt.Println("------- Pulling user comments")

	lastCommentID := ""
	commentIds := make([]string, 0)
	commentOptions := &reddit.ListUserOverviewOptions{
		ListOptions: reddit.ListOptions{
			Limit: LIMIT,
		},
		Sort: "new",
		Time: "all",
	}
	for {
		if lastCommentID != "" {
			commentOptions.ListOptions.After = lastCommentID
		}

		comments, _, err := userService.Comments(ctx, commentOptions)
		if err != nil {
			log.Fatal(err)
		}

		if len(comments) == 0 {
			break
		}

		for _, c := range comments {
			if c.Created.Time.Before(THRESHOLD) ||
				c.Created.Time.Equal(THRESHOLD) {
				commentIds = append(commentIds, c.FullID)
			}
		}

		lastCommentID = comments[len(comments)-1].FullID
	}

	// Delete all comments
	fmt.Println("------- Deleting user comments:", len(commentIds))

	for _, cID := range commentIds {
		fmt.Println(cID)
		//_, err := commentService.Delete(ctx, cID)
		//if err != nil {
		//	log.Fatal(err)
		//}
	}
}
