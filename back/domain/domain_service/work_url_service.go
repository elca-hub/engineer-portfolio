package domain_service

import (
	"devport/domain/repo/file_storage"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type WorkUrlService struct {
}

func NewWorkUrlService(workImageStorage file_storage.WorkImageStorageRepository) *WorkUrlService {
	return &WorkUrlService{}
}

func (s *WorkUrlService) extractTitle(body io.Reader) (string, error) {
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return "", fmt.Errorf("title tag not found")
		case html.StartTagToken:
			t := z.Token()
			if t.Data == "title" {
				z.Next()
				return strings.TrimSpace(z.Token().Data), nil
			}
		}
	}
}

func (s *WorkUrlService) FetchTitle(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTPリクエストエラー:", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("HTTPステータスエラー:", resp.Status)
		return "", err
	}

	title, err := s.extractTitle(resp.Body)
	if err != nil {
		fmt.Println("タイトル抽出エラー:", err)
		return "", err
	}

	return title, nil
}
