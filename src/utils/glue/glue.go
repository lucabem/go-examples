package GlueUtils

import (
	"context"

	AWSUtils "myApp/src/utils/aws"

	"github.com/aws/aws-sdk-go-v2/service/glue"
)

func GetColumnsFromGlueTable(database string, table string) ([]string, error) {

	sdkConfig, err := AWSUtils.GetSDKConfig()

	glueClient := glue.NewFromConfig(sdkConfig)
	out, err := glueClient.GetTable(context.TODO(), &glue.GetTableInput{
		DatabaseName: &database,
		Name:         &table,
	})

	if err != nil {
		return nil, err
	}

	var cols []string
	for _, ele := range out.Table.StorageDescriptor.Columns {
		cols = append(cols, *ele.Name)
	}

	return cols, nil
}
