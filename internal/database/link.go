package database

import (
	"errors"
	"time"
)

func parseTimestamp(unixTimestamp int64) string {
	time := time.Unix(unixTimestamp, 0)
	return time.Format("2006-01-02 15:04")
}

func LinkReadUri(alias string) (string, error) {
	var result UriLinkEntry

	dbResponse := DB.Find(&result, UriLinkEntry{Alias: alias})

	if dbResponse.Error != nil {
		return "", errors.New(dbResponse.Error.Error())
	}

	LinkRecordActivation(result.ID, false)

	return result.Url, nil
}

func LinkReadMetadata(alias string) (string, string, int64, int64, error) {
	var metadataResult UriLinkEntry

	dbResponseMetadata := DB.Find(&metadataResult, UriLinkEntry{Alias: alias})

	if dbResponseMetadata.Error != nil {
		return "", "", 0, 0, errors.New(dbResponseMetadata.Error.Error())
	}

	LinkRecordActivation(metadataResult.ID, true)

	var lookupCount int64
	DB.Model(&UriActivationEvent{}).Where(&UriActivationEvent{
		UriLinkEntryId: metadataResult.ID,
		IsRedirect:     false,
	}).Count(&lookupCount)

	var redirectCount int64
	DB.Model(&UriActivationEvent{}).Where(&UriActivationEvent{
		UriLinkEntryId: metadataResult.ID,
		IsRedirect:     true,
	}).Count(&redirectCount)

	createdAtStr := parseTimestamp(metadataResult.CreatedAt)
	return metadataResult.Url, createdAtStr, lookupCount, redirectCount, nil
}

func LinkRecordActivation(id uint, isRedirect bool) error {
	activation := UriActivationEvent{
		UriLinkEntryId: id,
		IsRedirect:     isRedirect,
	}

	dbResponse := DB.Create(&activation)

	if dbResponse.Error != nil {
		return errors.New(dbResponse.Error.Error())
	}

	return nil
}

func LinkCreateNew(url string, alias string) (string, string, error) {
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
