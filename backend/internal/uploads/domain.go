package uploads

import "time"

// Upload is the file metadata record stored in the database.
type Upload struct {
	ID           string    `json:"id"`
	UserID       *string   `json:"user_id,omitempty"`
	Filename     string    `json:"filename"`
	OriginalName string    `json:"original_name"`
	MimeType     string    `json:"mime_type"`
	Size         int64     `json:"size"`
	Provider     string    `json:"provider"` // local|s3
	Path         string    `json:"path"`
	URL          string    `json:"url"`
	CreatedAt    time.Time `json:"created_at"`
}
