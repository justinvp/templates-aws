using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleBucket = new Aws.S3.Bucket("exampleBucket", new Aws.S3.BucketArgs
        {
        });
        var exampleBucketPublicAccessBlock = new Aws.S3.BucketPublicAccessBlock("exampleBucketPublicAccessBlock", new Aws.S3.BucketPublicAccessBlockArgs
        {
            BlockPublicAcls = true,
            BlockPublicPolicy = true,
            Bucket = exampleBucket.Id,
        });
    }

}

