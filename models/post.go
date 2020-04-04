package models

// Post ...
// main model for a log entry
type Post struct {
	ID          uint32 `json:"id"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	TimeCreated uint64 `json:"time_created"`
}
