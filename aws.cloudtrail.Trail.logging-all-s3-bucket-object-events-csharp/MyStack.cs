using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
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
                                "arn:aws:s3:::",
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

