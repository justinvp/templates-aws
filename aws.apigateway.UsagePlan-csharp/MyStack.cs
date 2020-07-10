using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var myapi = new Aws.ApiGateway.RestApi("myapi", new Aws.ApiGateway.RestApiArgs
        {
        });
        var dev = new Aws.ApiGateway.Deployment("dev", new Aws.ApiGateway.DeploymentArgs
        {
            RestApi = myapi.Id,
            StageName = "dev",
        });
        var prod = new Aws.ApiGateway.Deployment("prod", new Aws.ApiGateway.DeploymentArgs
        {
            RestApi = myapi.Id,
            StageName = "prod",
        });
        var myUsagePlan = new Aws.ApiGateway.UsagePlan("myUsagePlan", new Aws.ApiGateway.UsagePlanArgs
        {
            ApiStages = 
            {
                new Aws.ApiGateway.Inputs.UsagePlanApiStageArgs
                {
                    ApiId = myapi.Id,
                    Stage = dev.StageName,
                },
                new Aws.ApiGateway.Inputs.UsagePlanApiStageArgs
                {
                    ApiId = myapi.Id,
                    Stage = prod.StageName,
                },
            },
            Description = "my description",
            ProductCode = "MYCODE",
            QuotaSettings = new Aws.ApiGateway.Inputs.UsagePlanQuotaSettingsArgs
            {
                Limit = 20,
                Offset = 2,
                Period = "WEEK",
            },
            ThrottleSettings = new Aws.ApiGateway.Inputs.UsagePlanThrottleSettingsArgs
            {
                BurstLimit = 5,
                RateLimit = 10,
            },
        });
    }

}

