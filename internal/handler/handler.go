package handler

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"github.com/candy12t/jarvis/internal/aws/service/s3"
	"github.com/candy12t/jarvis/internal/config"
	"github.com/candy12t/jarvis/internal/fetcher"
	"github.com/candy12t/jarvis/internal/fetcher/justwatch"
	"github.com/candy12t/jarvis/internal/model"
	"github.com/candy12t/jarvis/internal/notifier"
)

const dateFormat = "2006-01-02"

func Apply() error {
	timeNow := time.Now().UTC().Format(dateFormat)

	// create s3 client
	s3Client := s3.NewS3Service(config.AWSS3Bucket(), config.AWSS3Key())

	// download contents.json from s3
	storedContents, err := s3Client.Download()
	if err != nil {
		return err
	}

	// bytes to model.WrappedContents
	var unmarshledContents model.WrappedContents
	if err := json.Unmarshal(storedContents, &unmarshledContents); err != nil {
		return err
	}

	// get contents with justwatch
	fetcherClient := justwatch.NewClient(config.JustwatchCountry())
	_fetcher := fetcher.NewClient(fetcherClient, timeNow)
	contents, err := _fetcher.FetchContents()
	if err != nil {
		return err
	}

	// model.WrappedContents to bytes
	result := model.WrappedContents{Contents: contents}
	marshledResult, err := json.Marshal(result)
	if err != nil {
		return err
	}

	// compare contents.json and justwatch contents -> filter
	filterdContents := filterNewContents(unmarshledContents, result)

	// notifiy contents
	_notifier := notifier.NewClient(NotifyClient(), timeNow)
	_notifier.Notify(filterdContents)

	// overwirte contents.json with justwatch contents
	if err := s3Client.Upload(strings.NewReader(string(marshledResult))); err != nil {
		return err
	}

	return nil
}

func filterNewContents(oldContents, newContents model.WrappedContents) model.Contents {
	if reflect.DeepEqual(oldContents.Contents, newContents.Contents) {
		return model.Contents{}
	}

	if len(oldContents.Contents) == 0 || len(newContents.Contents) == 0 {
		return newContents.Contents
	}

	contents := make(model.Contents, 0, 0)

	for _, newContent := range newContents.Contents {
		flag := false
		for _, oldContent := range oldContents.Contents {
			if newContent == oldContent {
				flag = true
			}
		}
		if !flag {
			contents = append(contents, newContent)
		}
	}

	return contents
}
