package main

import (
	"context"
	"fmt"
)

func handlerRss(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"

	rss, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", rss)
	return nil
}
