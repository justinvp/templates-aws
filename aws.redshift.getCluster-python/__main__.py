import pulumi
import pulumi_aws as aws

test_cluster = aws.redshift.get_cluster(cluster_identifier="test-cluster")
test_stream = aws.kinesis.FirehoseDeliveryStream("testStream",
    destination="redshift",
    redshift_configuration={
        "clusterJdbcurl": f"jdbc:redshift://{test_cluster.endpoint}/{test_cluster.database_name}",
        "copyOptions": "delimiter '|'",
        "dataTableColumns": "test-col",
        "dataTableName": "test-table",
        "password": "T3stPass",
        "role_arn": aws_iam_role["firehose_role"]["arn"],
        "username": "testuser",
    },
    s3_configuration={
        "bucketArn": aws_s3_bucket["bucket"]["arn"],
        "bufferInterval": 400,
        "bufferSize": 10,
        "compressionFormat": "GZIP",
        "role_arn": aws_iam_role["firehose_role"]["arn"],
    })

