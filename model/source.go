package model

type Source struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

func (s Source) GetID() int64 {
	return s.ID
}
