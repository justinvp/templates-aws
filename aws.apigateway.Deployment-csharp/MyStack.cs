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
            RestApi = myDemoAPI.Id,
            ParentId = myDemoAPI.RootResourceId,
            PathPart = "test",
        });
        var myDemoMethod = new Aws.ApiGateway.Method("myDemoMethod", new Aws.ApiGateway.MethodArgs
        {
            RestApi = myDemoAPI.Id,
            ResourceId = myDemoResource.Id,
            HttpMethod = "GET",
            Authorization = "NONE",
        });
        var myDemoIntegration = new Aws.ApiGateway.Integration("myDemoIntegration", new Aws.ApiGateway.IntegrationArgs
        {
            RestApi = myDemoAPI.Id,
            ResourceId = myDemoResource.Id,
            HttpMethod = myDemoMethod.HttpMethod,
            Type = "MOCK",
        });
        var myDemoDeployment = new Aws.ApiGateway.Deployment("myDemoDeployment", new Aws.ApiGateway.DeploymentArgs
        {
            RestApi = myDemoAPI.Id,
            StageName = "test",
            Variables = 
            {
                { "answer", "42" },
            },
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                myDemoIntegration,
            },
        });
    }

}

