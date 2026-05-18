package main

import (
	"ZafirChowdhury/gatorGo/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("Invalid usage. Usage: %s <name> <url>", cmd.Name)
	}

	name, url := cmd.Args[0], cmd.Args[1]
	currentUser, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    currentUser.ID,
	})

	fmt.Println("Feed craeted sucesfully:")
	printFeed(feed, currentUser.Name)
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
