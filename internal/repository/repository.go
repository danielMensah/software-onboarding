package repository

type Story struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
	By string `json:"by"`
}

type Job struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Text string `json:"text"`
	Url string `json:"url"`
	By string `json:"by"`
}

type Repository interface {
	GetItems(table string, index string, items interface{}) error
	SaveItem(table string, item interface{}) error
}