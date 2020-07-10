using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var current = Output.Create(Aws.GetCallerIdentity.InvokeAsync());
        var foo = new Aws.S3.Bucket("foo", new Aws.S3.BucketArgs
        {
            ForceDestroy = true,
            Policy = current.Apply(current => @$"{{
    ""Version"": ""2012-10-17"",
    ""Statement"": [
        {{
            ""Sid"": ""AWSCloudTrailAclCheck"",
            ""Effect"": ""Allow"",
            ""Principal"": {{
              ""Service"": ""cloudtrail.amazonaws.com""
            }},
            ""Action"": ""s3:GetBucketAcl"",
            ""Resource"": ""arn:aws:s3:::tf-test-trail""
        }},
        {{
            ""Sid"": ""AWSCloudTrailWrite"",
            ""Effect"": ""Allow"",
            ""Principal"": {{
              ""Service"": ""cloudtrail.amazonaws.com""
            }},
            ""Action"": ""s3:PutObject"",
            ""Resource"": ""arn:aws:s3:::tf-test-trail/prefix/AWSLogs/{current.AccountId}/*"",
            ""Condition"": {{
                ""StringEquals"": {{
                    ""s3:x-amz-acl"": ""bucket-owner-full-control""
                }}
            }}
        }}
    ]
}}

"),
        });
        var foobar = new Aws.CloudTrail.Trail("foobar", new Aws.CloudTrail.TrailArgs
        {
            IncludeGlobalServiceEvents = false,
            S3BucketName = foo.Id,
            S3KeyPrefix = "prefix",
        });
    }

}

