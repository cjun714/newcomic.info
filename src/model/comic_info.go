package model

type ComicInfo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CoverURL    string `json:"cover_url"`
	DownloadURL string `json:"download_url"`
	Tags        string `json:"tags"`
	Year        int    `json:"year"`
	Pages       int    `json:"pages"`
	Size        string `json:"size"`
	Download    bool   `json:"download"`
	Favorite    bool   `json:"favorite"`
}
