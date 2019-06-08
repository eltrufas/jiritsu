package model

import "encoding/json"

type PageQuery struct {
	PageSize  int64
	PageToken int64
}

type Page struct {
	Data          []Model
	nextPageToken *int64
}

func (p *Page) NextPageToken() int64 {
	if p.nextPageToken != nil {
		return *p.nextPageToken
	}

	var token int64
	for _, m := range p.Data {
		id := m.GetID()
		if id > token {
			token = id
		}
	}
	p.nextPageToken = &token
	return token
}

func (p *Page) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Data          []Model `json:"data"`
		NextPageToken int64   `json:"next_page_token"`
	}{
		Data:          p.Data,
		NextPageToken: p.NextPageToken(),
	})
}
