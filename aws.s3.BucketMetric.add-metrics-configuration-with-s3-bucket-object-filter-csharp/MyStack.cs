using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.S3.Bucket("example", new Aws.S3.BucketArgs
        {
        });
        var example_filtered = new Aws.S3.BucketMetric("example-filtered", new Aws.S3.BucketMetricArgs
        {
            Bucket = example.BucketName,
            Filter = new Aws.S3.Inputs.BucketMetricFilterArgs
            {
                Prefix = "documents/",
                Tags = 
                {
                    { "class", "blue" },
                    { "priority", "high" },
                },
            },
        });
    }

}

