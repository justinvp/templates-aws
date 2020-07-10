using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var selectedRestApi = Output.Create(Aws.ApiGateway.GetRestApi.InvokeAsync(new Aws.ApiGateway.GetRestApiArgs
        {
            Name = @var.Api_gateway_name,
        }));
        var selectedUserPools = Output.Create(Aws.Cognito.GetUserPools.InvokeAsync(new Aws.Cognito.GetUserPoolsArgs
        {
            Name = @var.Cognito_user_pool_name,
        }));
        var cognito = new Aws.ApiGateway.Authorizer("cognito", new Aws.ApiGateway.AuthorizerArgs
        {
            ProviderArns = selectedUserPools.Apply(selectedUserPools => selectedUserPools.Arns),
            RestApi = selectedRestApi.Apply(selectedRestApi => selectedRestApi.Id),
            Type = "COGNITO_USER_POOLS",
        });
    }

}

