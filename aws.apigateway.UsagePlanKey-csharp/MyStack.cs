using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = new Aws.ApiGateway.RestApi("test", new Aws.ApiGateway.RestApiArgs
        {
        });
        var myusageplan = new Aws.ApiGateway.UsagePlan("myusageplan", new Aws.ApiGateway.UsagePlanArgs
        {
            ApiStages = 
            {
                new Aws.ApiGateway.Inputs.UsagePlanApiStageArgs
                {
                    ApiId = test.Id,
                    Stage = aws_api_gateway_deployment.Foo.Stage_name,
                },
            },
        });
        var mykey = new Aws.ApiGateway.ApiKey("mykey", new Aws.ApiGateway.ApiKeyArgs
        {
        });
        var main = new Aws.ApiGateway.UsagePlanKey("main", new Aws.ApiGateway.UsagePlanKeyArgs
        {
            KeyId = mykey.Id,
            KeyType = "API_KEY",
            UsagePlanId = myusageplan.Id,
        });
    }

}

