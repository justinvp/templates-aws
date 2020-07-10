using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.S3.Bucket("example", new Aws.S3.BucketArgs
        {
        });
        var example_filtered = new Aws.S3.AnalyticsConfiguration("example-filtered", new Aws.S3.AnalyticsConfigurationArgs
        {
            Bucket = example.BucketName,
            Filter = new Aws.S3.Inputs.AnalyticsConfigurationFilterArgs
            {
                Prefix = "documents/",
                Tags = 
                {
                    { "priority", "high" },
                    { "class", "blue" },
                },
            },
        });
    }

}

