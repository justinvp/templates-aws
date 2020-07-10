using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var hogeBucket = new Aws.S3.Bucket("hogeBucket", new Aws.S3.BucketArgs
        {
            Region = "us-east-1",
        });
        var hogeBucketPolicy = new Aws.S3.BucketPolicy("hogeBucketPolicy", new Aws.S3.BucketPolicyArgs
        {
            Bucket = hogeBucket.BucketName,
            Policy = @"{
    ""Version"": ""2012-10-17"",
    ""Statement"": [
        {
            ""Sid"": ""SSMBucketPermissionsCheck"",
            ""Effect"": ""Allow"",
            ""Principal"": {
                ""Service"": ""ssm.amazonaws.com""
            },
            ""Action"": ""s3:GetBucketAcl"",
            ""Resource"": ""arn:aws:s3:::tf-test-bucket-1234""
        },
        {
            ""Sid"": "" SSMBucketDelivery"",
            ""Effect"": ""Allow"",
            ""Principal"": {
                ""Service"": ""ssm.amazonaws.com""
            },
            ""Action"": ""s3:PutObject"",
            ""Resource"": [""arn:aws:s3:::tf-test-bucket-1234/*""],
            ""Condition"": {
                ""StringEquals"": {
                    ""s3:x-amz-acl"": ""bucket-owner-full-control""
                }
            }
        }
    ]
}

",
        });
        var foo = new Aws.Ssm.ResourceDataSync("foo", new Aws.Ssm.ResourceDataSyncArgs
        {
            S3Destination = new Aws.Ssm.Inputs.ResourceDataSyncS3DestinationArgs
            {
                BucketName = hogeBucket.BucketName,
                Region = hogeBucket.Region,
            },
        });
    }

}

