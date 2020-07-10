using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.Iam.OpenIdConnectProvider("default", new Aws.Iam.OpenIdConnectProviderArgs
        {
            ClientIdLists = 
            {
                "266362248691-342342xasdasdasda-apps.googleusercontent.com",
            },
            ThumbprintLists = {},
            Url = "https://accounts.google.com",
        });
    }

}

