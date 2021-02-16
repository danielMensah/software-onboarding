package hackernews

import repo "github.com/gymshark/software-onboarding/internal/repository"

const (
	Stories = "https://hacker-news.firebaseio.com/v0/item/8863.json"
	Jobs    = "https://hacker-news.firebaseio.com/v0/item/192327.json"
)

type HackerNewsAPI struct {
}

func New() *HackerNewsAPI {
	panic("implement me")
}

func (h *HackerNewsAPI) GetStories() []repo.Story {
	panic("implement me")
}

func (h *HackerNewsAPI) GetJobs() []repo.Job {
	panic("implement me")
}
