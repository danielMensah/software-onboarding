package hackernews

import (
	"fmt"
	repo "github.com/gymshark/software-onboarding/internal/repository"
	"github.com/gymshark/software-onboarding/internal/request"
	"github.com/sirupsen/logrus"
)

const (
	TopStoriesAPI = "https://hacker-news.firebaseio.com/v0/topstories.json"
	ItemAPI       = "https://hacker-news.firebaseio.com/v0/item/"
)

func GetStories() ([]repo.Story, error) {
	var storiesIds []int
	var stories []repo.Story

	err := request.Get(TopStoriesAPI, &storiesIds)

	if err != nil {
		return nil, err
	}

	// Make it a goroutine
	for _, id := range storiesIds {
		var story repo.Story

		err = request.Get(getItemUrl(id), &story)

		if err != nil {
			logrus.WithError(err).Error("error fetching story")
		} else {
			stories = append(stories, story)
		}
	}

	return stories, nil
}

func GetJobs() ([]repo.Job, error) {
	panic("implement me")
}

func getItemUrl(id int) string {
	return fmt.Sprintf("%s%d.json", ItemAPI, id)
}
