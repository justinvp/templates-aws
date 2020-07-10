using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testCluster = Output.Create(Aws.RedShift.GetCluster.InvokeAsync(new Aws.RedShift.GetClusterArgs
        {
            ClusterIdentifier = "test-cluster",
        }));
        var testStream = new Aws.Kinesis.FirehoseDeliveryStream("testStream", new Aws.Kinesis.FirehoseDeliveryStreamArgs
        {
            Destination = "redshift",
            RedshiftConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamRedshiftConfigurationArgs
            {
                ClusterJdbcurl = Output.Tuple(testCluster, testCluster).Apply(values =>
                {
                    var testCluster = values.Item1;
                    var testCluster1 = values.Item2;
                    return $"jdbc:redshift://{testCluster.Endpoint}/{testCluster1.DatabaseName}";
                }),
                CopyOptions = "delimiter '|'",
                DataTableColumns = "test-col",
                DataTableName = "test-table",
                Password = "T3stPass",
                RoleArn = aws_iam_role.Firehose_role.Arn,
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

