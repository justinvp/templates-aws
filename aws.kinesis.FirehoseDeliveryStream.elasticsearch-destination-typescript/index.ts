import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testCluster = new aws.elasticsearch.Domain("test_cluster", {});
const testStream = new aws.kinesis.FirehoseDeliveryStream("test_stream", {
    destination: "elasticsearch",
    elasticsearchConfiguration: {
        domainArn: testCluster.arn,
        indexName: "test",
        processingConfiguration: {
            enabled: true,
            processors: [{
                parameters: [{
                    parameterName: "LambdaArn",
                    parameterValue: pulumi.interpolate`${aws_lambda_function_lambda_processor.arn}:$LATEST`,
                }],
                type: "Lambda",
            }],
        },
        roleArn: aws_iam_role_firehose_role.arn,
        typeName: "test",
    },
    s3Configuration: {
        bucketArn: aws_s3_bucket_bucket.arn,
        bufferInterval: 400,
        bufferSize: 10,
        compressionFormat: "GZIP",
        roleArn: aws_iam_role_firehose_role.arn,
    },
});

