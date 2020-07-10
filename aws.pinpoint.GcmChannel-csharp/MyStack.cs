using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var app = new Aws.Pinpoint.App("app", new Aws.Pinpoint.AppArgs
        {
        });
        var gcm = new Aws.Pinpoint.GcmChannel("gcm", new Aws.Pinpoint.GcmChannelArgs
        {
            ApiKey = "api_key",
            ApplicationId = app.ApplicationId,
        });
    }

}

