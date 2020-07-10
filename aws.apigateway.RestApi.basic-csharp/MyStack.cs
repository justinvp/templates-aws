using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var myDemoAPI = new Aws.ApiGateway.RestApi("myDemoAPI", new Aws.ApiGateway.RestApiArgs
        {
            Description = "This is my API for demonstration purposes",
        });
    }

}

