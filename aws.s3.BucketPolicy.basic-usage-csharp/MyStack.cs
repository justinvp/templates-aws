using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
        });
        var bucketPolicy = new Aws.S3.BucketPolicy("bucketPolicy", new Aws.S3.BucketPolicyArgs
        {
            Bucket = bucket.Id,
            Policy = @"{
  ""Version"": ""2012-10-17"",
  ""Id"": ""MYBUCKETPOLICY"",
  ""Statement"": [
    {
      ""Sid"": ""IPAllow"",
      ""Effect"": ""Deny"",
      ""Principal"": ""*"",
      ""Action"": ""s3:*"",
      ""Resource"": ""arn:aws:s3:::my_tf_test_bucket/*"",
      ""Condition"": {
         ""IpAddress"": {""aws:SourceIp"": ""8.8.8.8/32""}
      }
    }
  ]
}

",
        });
    }

}

