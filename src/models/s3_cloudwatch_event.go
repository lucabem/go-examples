package Models

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type S3ObjectDetail struct {
	Version string `json:"version"`
	Bucket  struct {
		Name string `json:"name"`
	} `json:"bucket"`
	Object struct {
		Key       string `json:"key"`
		Size      int    `json:"size"`
		ETag      string `json:"eTag"`
		VersionId string `json:"version-id"`
		Sequencer string `json:"sequencer"`
	} `json:"object"`
	RequestID string `json:"request-id"`
	Requester string `json:"requester"`
	SourceIP  string `json:"source-ip-address"`
	Reason    string `json:"reason"`
}

func ParseCloudWatchEvent(request events.CloudWatchEvent) (S3ObjectDetail, error) {
	var eventDetail S3ObjectDetail
	err := json.Unmarshal(request.Detail, &eventDetail)
	if err != nil {
		return eventDetail, err
	}

	return eventDetail, nil

}
