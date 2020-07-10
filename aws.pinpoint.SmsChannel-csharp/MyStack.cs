using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var app = new Aws.Pinpoint.App("app", new Aws.Pinpoint.AppArgs
        {
        });
        var sms = new Aws.Pinpoint.SmsChannel("sms", new Aws.Pinpoint.SmsChannelArgs
        {
            ApplicationId = app.ApplicationId,
        });
    }

}

