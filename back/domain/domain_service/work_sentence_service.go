package domain_service

import (
	"devport/domain/repo/file_storage"
	"fmt"
	"regexp"
	"strings"
)

type WorkSentenceService struct {
	workImageStorage file_storage.WorkImageStorageRepository

	imageLen int
}

const (
	MAX_SENTENCE_IMAGE_LEN = 20
)

func NewWorkSentenceService(workImageStorage file_storage.WorkImageStorageRepository) *WorkSentenceService {
	return &WorkSentenceService{
		workImageStorage: workImageStorage,
	}
}

func (s *WorkSentenceService) GetImageIds(userId string, workId string, content string) []string {
	imageIds := []string{}
	regStr := fmt.Sprintf(`!\[.*\]\(%s/[a-z0-9-]{36}\.[a-z]{3,4}\)`, s.workImageStorage.GetPublicDomain(userId, workId))
	re := regexp.MustCompile(regStr)
	matches := re.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		// 後ろから/を探して、その前の文字列を取得
		url := match[0]
		lastSlashIndex := strings.LastIndex(url, "/")
		if lastSlashIndex != -1 {
			imageIds = append(imageIds, url[lastSlashIndex+1:len(url)-1])
		}
	}

	s.imageLen = len(imageIds)

	return workUnique(imageIds)
}

func (s *WorkSentenceService) IsFullImage() bool {
	return s.imageLen > MAX_SENTENCE_IMAGE_LEN
}

func workUnique(slice []string) []string {
	keys := make(map[string]struct{})
	result := []string{}
	for _, v := range slice {
		if _, ok := keys[v]; !ok {
			keys[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}
