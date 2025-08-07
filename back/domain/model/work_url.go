package model

import (
	"errors"
	"strings"
)

type WorkUrl struct {
	id    string
	url   string
	title string
}

func updateWorkUrlLinkLogic(url string) (string, error) {
	if !strings.HasPrefix(url, "https://") {
		return "", errors.New("URLはhttpsで始まる必要があります")
	}
	return url, nil
}

func NewWorkUrl(id string, url string, title string) (*WorkUrl, error) {
	if _, err := updateWorkUrlLinkLogic(url); err != nil {
		return nil, err
	}

	return &WorkUrl{id: id, url: url, title: title}, nil
}
