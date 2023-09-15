package datevapi

import (
	"io"
	"net/url"
)

type Form interface {
	Values() url.Values
	Files() map[string]FormFile
}

type FormFile struct {
	Filename string
	Content  io.Reader
}

type FileMetaData struct {
	Category string `json:"category"`
	Folder   string `json:"folder"`
	Register string `json:"register"`
}
