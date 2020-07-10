using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Acl = "private",
        });
        var firehoseRole = new Aws.Iam.Role("firehoseRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Action"": ""sts:AssumeRole"",
      ""Principal"": {
        ""Service"": ""firehose.amazonaws.com""
      },
      ""Effect"": ""Allow"",
      ""Sid"": """"
    }
  ]
}

",
        });
        var testStream = new Aws.Kinesis.FirehoseDeliveryStream("testStream", new Aws.Kinesis.FirehoseDeliveryStreamArgs
        {
            Destination = "s3",
            S3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamS3ConfigurationArgs
            {
                BucketArn = bucket.Arn,
                RoleArn = firehoseRole.Arn,
            },
        });
    }

}

