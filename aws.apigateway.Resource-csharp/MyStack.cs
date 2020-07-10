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
        var myDemoResource = new Aws.ApiGateway.Resource("myDemoResource", new Aws.ApiGateway.ResourceArgs
        {
            ParentId = myDemoAPI.RootResourceId,
            PathPart = "mydemoresource",
            RestApi = myDemoAPI.Id,
        });
    }

}

