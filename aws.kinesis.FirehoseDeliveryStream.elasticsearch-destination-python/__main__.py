import pulumi
import pulumi_aws as aws

test_cluster = aws.elasticsearch.Domain("testCluster")
test_stream = aws.kinesis.FirehoseDeliveryStream("testStream",
    destination="elasticsearch",
    elasticsearch_configuration={
        "domainArn": test_cluster.arn,
        "indexName": "test",
        "processingConfiguration": {
            "enabled": "true",
            "processors": [{
                "parameters": [{
                    "parameterName": "LambdaArn",
                    "parameterValue": f"{aws_lambda_function['lambda_processor']['arn']}:$LATEST",
                }],
                "type": "Lambda",
            }],
        },
        "role_arn": aws_iam_role["firehose_role"]["arn"],
        "typeName": "test",
    },
    s3_configuration={
        "bucketArn": aws_s3_bucket["bucket"]["arn"],
        "bufferInterval": 400,
        "bufferSize": 10,
        "compressionFormat": "GZIP",
        "role_arn": aws_iam_role["firehose_role"]["arn"],
    })

