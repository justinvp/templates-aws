using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleBucket = new Aws.S3.Bucket("exampleBucket", new Aws.S3.BucketArgs
        {
        });
        var exampleAccessPoint = new Aws.S3.AccessPoint("exampleAccessPoint", new Aws.S3.AccessPointArgs
        {
            Bucket = exampleBucket.Id,
        });
    }

}

