package repository

import (
	"encoding/json"
	"strconv"
	"time"
)

type ItemType string

var (
	ItemTypeStory ItemType = "story"
	ItemTypeJob   ItemType = "job"
)

type Item struct {
	ID    string    `json:"id"`
	Title string    `json:"title"`
	Text  string    `json:"text"`
	Type  ItemType  `json:"type"`
	Time  time.Time `json:"time"`
	URL   string    `json:"url"`
	By    string    `json:"by"`
}

func (i *Item) UnmarshalJSON(data []byte) error {
	type Alias Item

	aux := &struct {
		ID int `json:"id"`
		Time int64 `json:"time"`
		*Alias
	}{
		Alias: (*Alias)(i),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	i.Time = time.Unix(aux.Time, 0)
	i.ID = strconv.Itoa(aux.ID)

	return nil
}

type Repository interface {
	GetItems(index string, items *[]Item) error
	SaveItems(items []Item) error
}
