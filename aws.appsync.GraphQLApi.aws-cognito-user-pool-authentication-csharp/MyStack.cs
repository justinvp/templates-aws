using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.AppSync.GraphQLApi("example", new Aws.AppSync.GraphQLApiArgs
        {
            AuthenticationType = "AMAZON_COGNITO_USER_POOLS",
            UserPoolConfig = new Aws.AppSync.Inputs.GraphQLApiUserPoolConfigArgs
            {
                AwsRegion = data.Aws_region.Current.Name,
                DefaultAction = "DENY",
                UserPoolId = aws_cognito_user_pool.Example.Id,
            },
        });
    }

}

