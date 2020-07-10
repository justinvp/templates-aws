using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var logBucket = new Aws.S3.Bucket("logBucket", new Aws.S3.BucketArgs
        {
            Acl = "log-delivery-write",
        });
        var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
        {
            Acl = "private",
            Loggings = 
            {
                new Aws.S3.Inputs.BucketLoggingArgs
                {
                    TargetBucket = logBucket.Id,
                    TargetPrefix = "log/",
                },
            },
        });
    }

}

