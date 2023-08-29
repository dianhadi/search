package entity

type Post struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	SeoTitle       string `json:"seo-title"`
	Content        string `json:"content"`
	PreviewContent string `json:"preview-content"`
	PublishDate    string `json:"publish-date"`
	AuthorID       string `json:"author-id"`
}
