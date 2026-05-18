package main

import (
	"ZafirChowdhury/gatorGo/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("Invalid usage. Usage: %s <name> <url>", cmd.Name)
	}

	name, url := cmd.Args[0], cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Feed craeted sucesfully:")
	printFeed(feed, user.Name)

	// add the feed to follow_feed
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Follwed %s sucesfully\n", url)

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	// weried ass code just to satisfy both me and the lesson requrments
	for _, feed := range feeds {
		printFeed(database.Feed{
			ID:        feed.ID,
			CreatedAt: feed.CreatedAt,
			UpdatedAt: feed.UpdatedAt,
			Name:      feed.Name,
			Url:       feed.Url,
		}, feed.UserName)
	}

	return nil
}

func printFeed(feed database.Feed, userName string) {
	fmt.Printf("* ID:             	 %s\n", feed.ID)
	fmt.Printf("* Created:        	 %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       	 %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          	 %s\n", feed.Name)
	fmt.Printf("* URL:           	 %s\n", feed.Url)
	fmt.Printf("* Created by:        %s\n", userName)
}
