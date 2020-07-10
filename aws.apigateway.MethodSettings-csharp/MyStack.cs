using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testRestApi = new Aws.ApiGateway.RestApi("testRestApi", new Aws.ApiGateway.RestApiArgs
        {
            Description = "This is my API for demonstration purposes",
        });
        var testDeployment = new Aws.ApiGateway.Deployment("testDeployment", new Aws.ApiGateway.DeploymentArgs
        {
            RestApi = testRestApi.Id,
            StageName = "dev",
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_api_gateway_integration.test",
            },
        });
        var testStage = new Aws.ApiGateway.Stage("testStage", new Aws.ApiGateway.StageArgs
        {
            Deployment = testDeployment.Id,
            RestApi = testRestApi.Id,
            StageName = "prod",
        });
        var testResource = new Aws.ApiGateway.Resource("testResource", new Aws.ApiGateway.ResourceArgs
        {
            ParentId = testRestApi.RootResourceId,
            PathPart = "mytestresource",
            RestApi = testRestApi.Id,
        });
        var testMethod = new Aws.ApiGateway.Method("testMethod", new Aws.ApiGateway.MethodArgs
        {
            Authorization = "NONE",
            HttpMethod = "GET",
            ResourceId = testResource.Id,
            RestApi = testRestApi.Id,
        });
        var methodSettings = new Aws.ApiGateway.MethodSettings("methodSettings", new Aws.ApiGateway.MethodSettingsArgs
        {
            MethodPath = Output.Tuple(testResource.PathPart, testMethod.HttpMethod).Apply(values =>
            {
                var pathPart = values.Item1;
                var httpMethod = values.Item2;
                return $"{pathPart}/{httpMethod}";
            }),
            RestApi = testRestApi.Id,
            Settings = new Aws.ApiGateway.Inputs.MethodSettingsSettingsArgs
            {
                LoggingLevel = "INFO",
                MetricsEnabled = true,
            },
            StageName = testStage.StageName,
        });
        var testIntegration = new Aws.ApiGateway.Integration("testIntegration", new Aws.ApiGateway.IntegrationArgs
        {
            HttpMethod = testMethod.HttpMethod,
            RequestTemplates = 
            {
                { "application/xml", @"{
   ""body"" : $input.json('$')
}

" },
            },
            ResourceId = testResource.Id,
            RestApi = testRestApi.Id,
            Type = "MOCK",
        });
    }

}

