package models

type DictionaryItem struct {
	Id          string `json:"id"`
	Term        string `json:"term"`
	Description string `json:"description"`
}
