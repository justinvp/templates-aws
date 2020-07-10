using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.Kinesis.VideoStream("default", new Aws.Kinesis.VideoStreamArgs
        {
            DataRetentionInHours = 1,
            DeviceName = "kinesis-video-device-name",
            MediaType = "video/h264",
            Tags = 
            {
                { "Name", "kinesis-video-stream" },
            },
        });
    }

}

