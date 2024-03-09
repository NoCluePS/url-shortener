package shortener

import (
	"fmt"
	"url-shortener/utils"
)

type URLShortener struct {
	Urls map[string]string
}

func (us *URLShortener) HandleShorten(originalUrl string) (string, error) {
	if originalUrl == "" {
		return "", fmt.Errorf("original url is required")
	}

	shortKey := utils.GenerateShortKey()
	us.Urls[shortKey] = originalUrl

	shortUrl := fmt.Sprintf("http://localhost:8080/short/%s", shortKey)

	return shortUrl, nil;
}

func (us *URLShortener) HandleRedirect(shortKey string) (string, error) {
    originalURL, ok := us.Urls[shortKey]
    if !ok {
		return "", fmt.Errorf("short url not found")
    }

	return originalURL, nil;
}