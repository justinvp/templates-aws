using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var gcmApplication = new Aws.Sns.PlatformApplication("gcmApplication", new Aws.Sns.PlatformApplicationArgs
        {
            Platform = "GCM",
            PlatformCredential = "<GCM API KEY>",
        });
    }

}

