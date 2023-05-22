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
