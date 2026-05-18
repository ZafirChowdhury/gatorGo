package main

import (
	"ZafirChowdhury/gatorGo/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Invalid command. Usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	// get feed
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error while trying to find feed. Error: %s", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
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
	fmt.Printf("* Feed: %s\n", feedFollow.FeedName)
	fmt.Printf("* Username: %s\n", feedFollow.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	follwing, err := s.db.GetFeedFollowsForUserByID(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, follow := range follwing {
		fmt.Printf("* %s\n", follow.FeedName)
	}

	return nil
}
