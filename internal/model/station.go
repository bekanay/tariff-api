package model

type Station struct {
	ID        int    `json:"id"`
	CODE      int    `json:"code"`
	NAME      string `json:"name"`
	PARAGRAPH string `json:"paragraph"`
}
