package entity

type Music struct {
	ID       string `json:"id"`
	Title    string `json:"title" binding:"required"`
	Artist   string `json:"artist"`
	URL      string `json:"url" binding:"required,url"`
	ThumbURL string `json:"thumb_url" binding:"omitempty,url"`
}
