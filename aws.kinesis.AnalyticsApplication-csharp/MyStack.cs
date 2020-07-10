using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testStream = new Aws.Kinesis.Stream("testStream", new Aws.Kinesis.StreamArgs
        {
            ShardCount = 1,
        });
        var testApplication = new Aws.Kinesis.AnalyticsApplication("testApplication", new Aws.Kinesis.AnalyticsApplicationArgs
        {
            Inputs = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsArgs
            {
                KinesisStream = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsKinesisStreamArgs
                {
                    ResourceArn = testStream.Arn,
                    RoleArn = aws_iam_role.Test.Arn,
                },
                NamePrefix = "test_prefix",
                Parallelism = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsParallelismArgs
                {
                    Count = 1,
                },
                Schema = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsSchemaArgs
                {
                    RecordColumns = 
                    {
                        new Aws.Kinesis.Inputs.AnalyticsApplicationInputsSchemaRecordColumnArgs
                        {
                            Mapping = "$.test",
                            Name = "test",
                            SqlType = "VARCHAR(8)",
                        },
                    },
                    RecordEncoding = "UTF-8",
                    RecordFormat = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsSchemaRecordFormatArgs
                    {
                        MappingParameters = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersArgs
                        {
                            Json = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersJsonArgs
                            {
                                RecordRowPath = "$",
                            },
                        },
                    },
                },
            },
        });
    }

}

