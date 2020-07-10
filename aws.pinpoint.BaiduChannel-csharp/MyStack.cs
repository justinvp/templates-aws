using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var app = new Aws.Pinpoint.App("app", new Aws.Pinpoint.AppArgs
        {
        });
        var channel = new Aws.Pinpoint.BaiduChannel("channel", new Aws.Pinpoint.BaiduChannelArgs
        {
            ApiKey = "",
            ApplicationId = app.ApplicationId,
            SecretKey = "",
        });
    }

}

