using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Acl = "private",
            LifecycleRules = 
            {
                new Aws.S3.Inputs.BucketLifecycleRuleArgs
                {
                    Enabled = true,
                    Expiration = new Aws.S3.Inputs.BucketLifecycleRuleExpirationArgs
                    {
                        Days = 90,
                    },
                    Id = "log",
                    Prefix = "log/",
                    Tags = 
                    {
                        { "autoclean", "true" },
                        { "rule", "log" },
                    },
                    Transition = 
                    {
                        
                        {
                            { "days", 30 },
                            { "storageClass", "STANDARD_IA" },
                        },
                        
                        {
                            { "days", 60 },
                            { "storageClass", "GLACIER" },
                        },
                    },
                },
                new Aws.S3.Inputs.BucketLifecycleRuleArgs
                {
                    Enabled = true,
                    Expiration = new Aws.S3.Inputs.BucketLifecycleRuleExpirationArgs
                    {
                        Date = "2016-01-12",
                    },
                    Id = "tmp",
                    Prefix = "tmp/",
                },
            },
        });
        var versioningBucket = new Aws.S3.Bucket("versioningBucket", new Aws.S3.BucketArgs
        {
            Acl = "private",
            LifecycleRules = 
            {
                new Aws.S3.Inputs.BucketLifecycleRuleArgs
                {
                    Enabled = true,
                    NoncurrentVersionExpiration = new Aws.S3.Inputs.BucketLifecycleRuleNoncurrentVersionExpirationArgs
                    {
                        Days = 90,
                    },
                    NoncurrentVersionTransition = 
                    {
                        
                        {
                            { "days", 30 },
                            { "storageClass", "STANDARD_IA" },
                        },
                        
                        {
                            { "days", 60 },
                            { "storageClass", "GLACIER" },
                        },
                    },
                    Prefix = "config/",
                },
            },
            Versioning = new Aws.S3.Inputs.BucketVersioningArgs
            {
                Enabled = true,
            },
        });
    }

}

