using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.S3.Bucket("example", new Aws.S3.BucketArgs
        {
        });
        var example_entire_bucket = new Aws.S3.BucketMetric("example-entire-bucket", new Aws.S3.BucketMetricArgs
        {
            Bucket = example.BucketName,
        });
    }

}

