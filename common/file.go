package common

// Metadata holds a bag of data representing information
// about an uploaded file
type Metadata struct {
	ID           string    `json:"id"`
	DefinitionID string    `json:"definitionID"`
	UserID       string    `json:"userID"`
	Header       string    `json:"header"`
	Data         []byte    `json:"-"`
	Path         string    `json:"path,omitempty"`
	Filename     string    `json:"filename"`
	SizeBytes    int64     `json:"sizeBytes"`
	MD5          string    `json:"hash"`
	CreatedAt    time.Time `json:"createdAt"`
}
