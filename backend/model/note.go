package model

import "time"

//Note store note information
type Note struct {
	//Line below which is called embed struct
	Model

	Title  string     `json:"title"`
	Body   string     `json:"body"`
	Tags   []string   `json:"tags"`
	Expire *time.Time `json:"expire"`
}
