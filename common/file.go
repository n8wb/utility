package common

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

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

// FilenameWithPath preserves any path filename that was uploaded
// with the file in multipath upload, we set the path as the file upload
// formfile, ie:
// -F key/path/blah=@file.ext
func (m *Metadata) FilenameWithPath() string {
	ext := filepath.Ext(m.Filename)
	pathname := strings.TrimSuffix(m.Filename, ext)

	if m.Path == m.Filename || m.Path == "" || m.Path == pathname {
		return m.Filename
	}

	return fmt.Sprintf("%s/%s", m.Path, m.Filename)
}
