using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var apnsApplication = new Aws.Sns.PlatformApplication("apnsApplication", new Aws.Sns.PlatformApplicationArgs
        {
            Platform = "APNS",
            PlatformCredential = "<APNS PRIVATE KEY>",
            PlatformPrincipal = "<APNS CERTIFICATE>",
        });
    }

}

