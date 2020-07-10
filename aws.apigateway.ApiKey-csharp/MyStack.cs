using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var myDemoApiKey = new Aws.ApiGateway.ApiKey("myDemoApiKey", new Aws.ApiGateway.ApiKeyArgs
        {
        });
    }

}

