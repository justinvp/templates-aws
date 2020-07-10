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
            HttpMethod = myDemoMethod.HttpMethod,
            ResourceId = myDemoResource.Id,
            RestApi = myDemoAPI.Id,
            Type = "MOCK",
        });
        var response200 = new Aws.ApiGateway.MethodResponse("response200", new Aws.ApiGateway.MethodResponseArgs
        {
            HttpMethod = myDemoMethod.HttpMethod,
            ResourceId = myDemoResource.Id,
            RestApi = myDemoAPI.Id,
            StatusCode = "200",
        });
        var myDemoIntegrationResponse = new Aws.ApiGateway.IntegrationResponse("myDemoIntegrationResponse", new Aws.ApiGateway.IntegrationResponseArgs
        {
            HttpMethod = myDemoMethod.HttpMethod,
            ResourceId = myDemoResource.Id,
            ResponseTemplates = 
            {
                { "application/xml", @"#set($inputRoot = $input.path('$'))
<?xml version=""1.0"" encoding=""UTF-8""?>
<message>
    $inputRoot.body
</message>

" },
            },
            RestApi = myDemoAPI.Id,
            StatusCode = response200.StatusCode,
        });
    }

}

