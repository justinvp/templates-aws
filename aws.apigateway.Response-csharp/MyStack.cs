using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var main = new Aws.ApiGateway.RestApi("main", new Aws.ApiGateway.RestApiArgs
        {
        });
        var test = new Aws.ApiGateway.Response("test", new Aws.ApiGateway.ResponseArgs
        {
            ResponseParameters = 
            {
                { "gatewayresponse.header.Authorization", "'Basic'" },
            },
            ResponseTemplates = 
            {
                { "application/json", "{'message':$context.error.messageString}" },
            },
            ResponseType = "UNAUTHORIZED",
            RestApiId = main.Id,
            StatusCode = "401",
        });
    }

}

