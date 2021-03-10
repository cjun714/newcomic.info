package model

type ComicInfo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Cover       string `json:"cover"`
	PageURL     string `json:"page_url"`
	DownloadURL string `json:"download_url"`
	Tags        string `json:"tags"`
	Publisher   string `json:"publisher"`
	Year        int    `json:"year"`
	Pages       int    `json:"pages"`
	Size        int    `json:"size"`
	Download    bool   `json:"download" gorm:"index"`
	Favorite    bool   `json:"favorite" gorm:"index"`
	Delete      bool   `json:"delete" gorm:"index"`
}
