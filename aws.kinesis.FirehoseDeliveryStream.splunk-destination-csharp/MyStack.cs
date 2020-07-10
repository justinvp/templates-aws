using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testStream = new Aws.Kinesis.FirehoseDeliveryStream("testStream", new Aws.Kinesis.FirehoseDeliveryStreamArgs
        {
            Destination = "splunk",
            S3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamS3ConfigurationArgs
            {
                BucketArn = aws_s3_bucket.Bucket.Arn,
                BufferInterval = 400,
                BufferSize = 10,
                CompressionFormat = "GZIP",
                RoleArn = aws_iam_role.Firehose.Arn,
            },
            SplunkConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamSplunkConfigurationArgs
            {
                HecAcknowledgmentTimeout = 600,
                HecEndpoint = "https://http-inputs-mydomain.splunkcloud.com:443",
                HecEndpointType = "Event",
                HecToken = "51D4DA16-C61B-4F5F-8EC7-ED4301342A4A",
                S3BackupMode = "FailedEventsOnly",
            },
        });
    }

}

