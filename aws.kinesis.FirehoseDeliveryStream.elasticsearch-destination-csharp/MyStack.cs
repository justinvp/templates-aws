using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testCluster = new Aws.ElasticSearch.Domain("testCluster", new Aws.ElasticSearch.DomainArgs
        {
        });
        var testStream = new Aws.Kinesis.FirehoseDeliveryStream("testStream", new Aws.Kinesis.FirehoseDeliveryStreamArgs
        {
            Destination = "elasticsearch",
            ElasticsearchConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamElasticsearchConfigurationArgs
            {
                DomainArn = testCluster.Arn,
                IndexName = "test",
                ProcessingConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationArgs
                {
                    Enabled = true,
                    Processors = 
                    {
                        new Aws.Kinesis.Inputs.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationProcessorArgs
                        {
                            Parameters = 
                            {
                                new Aws.Kinesis.Inputs.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationProcessorParameterArgs
                                {
                                    ParameterName = "LambdaArn",
                                    ParameterValue = $"{aws_lambda_function.Lambda_processor.Arn}:$LATEST",
                                },
                            },
                            Type = "Lambda",
                        },
                    },
                },
                RoleArn = aws_iam_role.Firehose_role.Arn,
                TypeName = "test",
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

