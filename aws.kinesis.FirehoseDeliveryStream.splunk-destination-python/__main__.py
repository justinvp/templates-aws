import pulumi
import pulumi_aws as aws

test_stream = aws.kinesis.FirehoseDeliveryStream("testStream",
    destination="splunk",
    s3_configuration={
        "bucketArn": aws_s3_bucket["bucket"]["arn"],
        "bufferInterval": 400,
        "bufferSize": 10,
        "compressionFormat": "GZIP",
        "role_arn": aws_iam_role["firehose"]["arn"],
    },
    splunk_configuration={
        "hecAcknowledgmentTimeout": 600,
        "hecEndpoint": "https://http-inputs-mydomain.splunkcloud.com:443",
        "hecEndpointType": "Event",
        "hecToken": "51D4DA16-C61B-4F5F-8EC7-ED4301342A4A",
        "s3BackupMode": "FailedEventsOnly",
    })

