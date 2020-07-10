using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var app = new Aws.Pinpoint.App("app", new Aws.Pinpoint.AppArgs
        {
        });
        var channel = new Aws.Pinpoint.AdmChannel("channel", new Aws.Pinpoint.AdmChannelArgs
        {
            ApplicationId = app.ApplicationId,
            ClientId = "",
            ClientSecret = "",
            Enabled = true,
        });
    }

}

