package database

import (
	"time"
)

func parseTimestamp(unixTimestamp int64) string {
	time := time.Unix(unixTimestamp, 0)
	return time.Format("2006-01-02 15:04")
}

func LinkRead(alias string) (string, string, error) {

	var result UriLinkEntry
	err := DB.Find(&result, UriLinkEntry{Alias: alias})

	if err == nil {
		return "", "", DB.Error
	}

	createdAtStr := parseTimestamp(result.CreatedAt)
	return result.Url, createdAtStr, nil
}

func LinkWrite(url string, alias string) (string, string, error) {
	link := UriLinkEntry{
		Url:       url,
		Alias:     alias,
		CreatedAt: 0,
	}

	result := DB.Create(&link)

	if result.Error != nil {
		return "", "", result.Error
	}

	createdAtStr := parseTimestamp(link.CreatedAt)
	return url, createdAtStr, nil
}
