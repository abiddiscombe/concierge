package database

import (
	"errors"
	"time"
)

func parseTimestamp(unixTimestamp int64) string {
	time := time.Unix(unixTimestamp, 0)
	return time.Format("2006-01-02 15:04")
}

func LinkRead(alias string) (string, string, error) {

	var result UriLinkEntry
	dbResponse := DB.Find(&result, UriLinkEntry{Alias: alias})

	if dbResponse.Error != nil {
		return "", "", errors.New(dbResponse.Error.Error())
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

	dbResponse := DB.Create(&link)

	if dbResponse.Error != nil {
		return "", "", errors.New(dbResponse.Error.Error())
	}

	createdAtStr := parseTimestamp(link.CreatedAt)
	return url, createdAtStr, nil
}
