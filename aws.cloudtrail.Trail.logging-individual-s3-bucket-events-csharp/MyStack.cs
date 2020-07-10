using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var important_bucket = Output.Create(Aws.S3.GetBucket.InvokeAsync(new Aws.S3.GetBucketArgs
        {
            Bucket = "important-bucket",
        }));
        var example = new Aws.CloudTrail.Trail("example", new Aws.CloudTrail.TrailArgs
        {
            EventSelectors = 
            {
                new Aws.CloudTrail.Inputs.TrailEventSelectorArgs
                {
                    DataResource = 
                    {
                        
                        {
                            { "type", "AWS::S3::Object" },
                            { "values", 
                            {
                                important_bucket.Apply(important_bucket => $"{important_bucket.Arn}/"),
                            } },
                        },
                    },
                    IncludeManagementEvents = true,
                    ReadWriteType = "All",
                },
            },
        });
    }

}

