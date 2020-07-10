using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testCluster = new Aws.RedShift.Cluster("testCluster", new Aws.RedShift.ClusterArgs
        {
            ClusterIdentifier = "tf-redshift-cluster-%d",
            ClusterType = "single-node",
            DatabaseName = "test",
            MasterPassword = "T3stPass",
            MasterUsername = "testuser",
            NodeType = "dc1.large",
        });
        var testStream = new Aws.Kinesis.FirehoseDeliveryStream("testStream", new Aws.Kinesis.FirehoseDeliveryStreamArgs
        {
            Destination = "redshift",
            RedshiftConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamRedshiftConfigurationArgs
            {
                ClusterJdbcurl = Output.Tuple(testCluster.Endpoint, testCluster.DatabaseName).Apply(values =>
                {
                    var endpoint = values.Item1;
                    var databaseName = values.Item2;
                    return $"jdbc:redshift://{endpoint}/{databaseName}";
                }),
                CopyOptions = "delimiter '|'",
                DataTableColumns = "test-col",
                DataTableName = "test-table",
                Password = "T3stPass",
                RoleArn = aws_iam_role.Firehose_role.Arn,
                S3BackupConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfigurationArgs
                {
                    BucketArn = aws_s3_bucket.Bucket.Arn,
                    BufferInterval = 300,
                    BufferSize = 15,
                    CompressionFormat = "GZIP",
                    RoleArn = aws_iam_role.Firehose_role.Arn,
                },
                S3BackupMode = "Enabled",
                Username = "testuser",
            },
            S3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamS3ConfigurationArgs
            {
                BucketArn = aws_s3_bucket.Bucket.Arn,
                BufferInterval = 400,
                BufferSize = 10,
                CompressionFormat = "GZIP",
                RoleArn = aws_iam_role.Firehose_role.Arn,
            },
        });
    }

}

