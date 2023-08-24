# XlsxToCsvTransformer
This repository contains the necessary code to convert an excel to csv transparently to the programmer, using CloudWatch events (putObject).
When an xlsx file is created in the specified path, it will trigger the lambda that will process the event and try to convert the file to CSV, using Glue to get the headers.

## How to run it using Docker
On docker's folder we have an Dockerfile and a docker-compose.yaml file that allows us to run it just with
```
docker-compose --file docker/docker-compose.yaml up
```
Notice that if you want to connect with your AWS account, you need to set up this environment variables en docker-compose file
```yaml
    environment:
      - AWS_REGION=eu-west-1
      - AWS_ACCESS_KEY_ID=VALUE
      - AWS_SECRET_ACCESS_KEY=VALUE
      - AWS_SESSION_TOKEN=VALUE
```
This statement will create an endpoint to `http://localhost:9000/2015-03-31/functions/function/invocations` that simulates an AWS Lambda function running on Go-1.21.1. For example, we could make requests to this endpoint such as 
```json
{
    "version": "0",
    "id": "example-id",
    "detail-type": "ObjectCreated",
    "source": "aws.s3",
    "account": "your-account-id",
    "time": "2023-08-23T12:34:56Z",
    "region": "us-east-1",
    "resources": [
        "arn:aws:s3:::my-example-bucket"
    ],
    "detail": {
        "version": "2.2",
        "bucket": {
            "name": "my-bucket"
        },
        "object": {
            "key": "path/to/my/file.xlsx",
            "size": 1024,
            "eTag": "example-etag",
            "versionId": "example-version-id",
            "sequencer": "example-sequencer"
        },
        "request-id": "IDD",
        "requester": "requester",
        "source-ip-address": "localhost",
        "reason": "reaso"
    }
}
```


## Generate ZIP Build for AWS Lambda
Inside cmd we have an script that generates ZIP file needed for run function on AWS Lambda. Remember that you need to give executable's permissions to this script
```sh
./cmd/run.sh
```
This script will genereate folder pkg/ which will contain the zip