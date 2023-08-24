package main

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	Models "myApp/src/models"
	CsvUtils "myApp/src/utils/csv"
	GlueUtils "myApp/src/utils/glue"
	S3Utils "myApp/src/utils/s3"
	StringsUtils "myApp/src/utils/strings"

	Transformations "myApp/src/transformations"
)

func handleRequest(request events.CloudWatchEvent) (string, error) {

	fmt.Printf("Data recived = \n%s\n", request)

	eventDetail, err := Models.ParseCloudWatchEvent(request)
	if err != nil {
		fmt.Printf("Couldn't format event. Is it an S3 CloudWatchEvent? %v\n", string(request.Detail))
		return "", err
	}

	bucket := eventDetail.Bucket.Name
	key := eventDetail.Object.Key

	obj, err := S3Utils.GetObject(bucket, key)

	if err != nil {
		fmt.Printf("Couldn't load key=%s in bucket=%s. Does it exist?\n", key, bucket)
		return "", err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(obj.Body)
	objReader := bytes.NewReader(buf.Bytes())

	var cols []string
	database := "my_glue_datatabase"
	table := "my_glue_table"
	cols, err = GlueUtils.GetColumnsFromGlueTable(database, table)
	if err != nil {
		fmt.Printf("Couldn't load table=%s in database=%s. Does it exist?\n", database, table)
		return "", err
	}

	hasHeader := true
	data, err := CsvUtils.ParseXLSXToCSV(objReader, cols, hasHeader)
	if err != nil {
		return "", err
	}

	err = CsvUtils.CreateCSVToFile(data, "/tmp/output.csv")

	if err != nil {
		return "", err
	}

	transformations := []func(string) string{
		Transformations.Trim,
		Transformations.ToLower,
		Transformations.ReplaceSpaces,
		Transformations.ReplaceExtension,
	}

	err = S3Utils.PutObject(
		"/tmp/output.csv",
		bucket,
		StringsUtils.TransformPath(key, transformations),
	)
	return "200", err
}

func main() {

	lambda.Start(handleRequest)
}
