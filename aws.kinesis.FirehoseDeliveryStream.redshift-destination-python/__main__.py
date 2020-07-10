import pulumi
import pulumi_aws as aws

test_cluster = aws.redshift.Cluster("testCluster",
    cluster_identifier="tf-redshift-cluster-%d",
    cluster_type="single-node",
    database_name="test",
    master_password="T3stPass",
    master_username="testuser",
    node_type="dc1.large")
test_stream = aws.kinesis.FirehoseDeliveryStream("testStream",
    destination="redshift",
    redshift_configuration={
        "clusterJdbcurl": pulumi.Output.all(test_cluster.endpoint, test_cluster.database_name).apply(lambda endpoint, database_name: f"jdbc:redshift://{endpoint}/{database_name}"),
        "copyOptions": "delimiter '|'",
        "dataTableColumns": "test-col",
        "dataTableName": "test-table",
        "password": "T3stPass",
        "role_arn": aws_iam_role["firehose_role"]["arn"],
        "s3BackupConfiguration": {
            "bucketArn": aws_s3_bucket["bucket"]["arn"],
            "bufferInterval": 300,
            "bufferSize": 15,
            "compressionFormat": "GZIP",
            "role_arn": aws_iam_role["firehose_role"]["arn"],
        },
        "s3BackupMode": "Enabled",
        "username": "testuser",
    },
    s3_configuration={
        "bucketArn": aws_s3_bucket["bucket"]["arn"],
        "bufferInterval": 400,
        "bufferSize": 10,
        "compressionFormat": "GZIP",
        "role_arn": aws_iam_role["firehose_role"]["arn"],
    })

