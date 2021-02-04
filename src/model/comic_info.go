package model

type ComicInfo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Cover       string `json:"cover"`
	URL         string `json:"url"`
	DownloadURL string `json:"download_url"`
	Tags        string `json:"tags"`
	Publisher   string `json:"publisher"`
	Year        int    `json:"year"`
	Pages       int    `json:"pages"`
	Size        int    `json:"size"`
	Download    bool   `json:"download"`
	Favorite    bool   `json:"favorite"`
}
