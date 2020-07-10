import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testStream = new aws.kinesis.FirehoseDeliveryStream("test_stream", {
    destination: "splunk",
    s3Configuration: {
        bucketArn: aws_s3_bucket_bucket.arn,
        bufferInterval: 400,
        bufferSize: 10,
        compressionFormat: "GZIP",
        roleArn: aws_iam_role_firehose.arn,
    },
    splunkConfiguration: {
        hecAcknowledgmentTimeout: 600,
        hecEndpoint: "https://http-inputs-mydomain.splunkcloud.com:443",
        hecEndpointType: "Event",
        hecToken: "51D4DA16-C61B-4F5F-8EC7-ED4301342A4A",
        s3BackupMode: "FailedEventsOnly",
    },
});

