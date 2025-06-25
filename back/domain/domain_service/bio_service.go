package domain_service

import (
	"devport/domain/repo/file_storage"
	"fmt"
	"regexp"
	"strings"
)

type BioService struct {
	bioImageStorage file_storage.BioImageStorageRepository

	imageLen int
}

const (
	MAX_IMAGE_LEN = 5
)

func NewBioService(bioImageStorage file_storage.BioImageStorageRepository) *BioService {
	return &BioService{
		bioImageStorage: bioImageStorage,
	}
}

func (s *BioService) GetImageIds(userId string, content string) []string {
	imageIds := []string{}
	regStr := fmt.Sprintf(`!\[.*\]\(%s/[a-z0-9-]{36}\.[a-z]{3,4}\)`, s.bioImageStorage.GetPublicDomain(userId))
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

	return unique(imageIds)
}

func (s *BioService) IsFullImage() bool {
	return s.imageLen > MAX_IMAGE_LEN
}

func unique(slice []string) []string {
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
