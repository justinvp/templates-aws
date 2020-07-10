using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Acl = "public-read",
            CorsRules = 
            {
                new Aws.S3.Inputs.BucketCorsRuleArgs
                {
                    AllowedHeaders = 
                    {
                        "*",
                    },
                    AllowedMethods = 
                    {
                        "PUT",
                        "POST",
                    },
                    AllowedOrigins = 
                    {
                        "https://s3-website-test.mydomain.com",
                    },
                    ExposeHeaders = 
                    {
                        "ETag",
                    },
                    MaxAgeSeconds = 3000,
                },
            },
        });
    }

}

