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
        var myDemoMethod = new Aws.ApiGateway.Method("myDemoMethod", new Aws.ApiGateway.MethodArgs
        {
            Authorization = "NONE",
            HttpMethod = "GET",
            ResourceId = myDemoResource.Id,
            RestApi = myDemoAPI.Id,
        });
        var myDemoIntegration = new Aws.ApiGateway.Integration("myDemoIntegration", new Aws.ApiGateway.IntegrationArgs
        {
            CacheKeyParameters = 
            {
                "method.request.path.param",
            },
            CacheNamespace = "foobar",
            HttpMethod = myDemoMethod.HttpMethod,
            RequestParameters = 
            {
                { "integration.request.header.X-Authorization", "'static'" },
            },
            RequestTemplates = 
            {
                { "application/xml", @"{
   ""body"" : $input.json('$')
}

" },
            },
            ResourceId = myDemoResource.Id,
            RestApi = myDemoAPI.Id,
            TimeoutMilliseconds = 29000,
            Type = "MOCK",
        });
    }

}

