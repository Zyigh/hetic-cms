package models

type PageForList struct {
	Name string `json:"name"`
}

type PagesList []PageForList

type SinglePage struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
