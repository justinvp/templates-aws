using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.DataSync.Task("example", new Aws.DataSync.TaskArgs
        {
            DestinationLocationArn = aws_datasync_location_s3.Destination.Arn,
            Options = new Aws.DataSync.Inputs.TaskOptionsArgs
            {
                BytesPerSecond = -1,
            },
            SourceLocationArn = aws_datasync_location_nfs.Source.Arn,
        });
    }

}

