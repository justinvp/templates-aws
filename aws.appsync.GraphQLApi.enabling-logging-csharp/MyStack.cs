using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleRole = new Aws.Iam.Role("exampleRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
    ""Version"": ""2012-10-17"",
    ""Statement"": [
        {
        ""Effect"": ""Allow"",
        ""Principal"": {
            ""Service"": ""appsync.amazonaws.com""
        },
        ""Action"": ""sts:AssumeRole""
        }
    ]
}

",
        });
        var exampleRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("exampleRolePolicyAttachment", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = "arn:aws:iam::aws:policy/service-role/AWSAppSyncPushToCloudWatchLogs",
            Role = exampleRole.Name,
        });
        var exampleGraphQLApi = new Aws.AppSync.GraphQLApi("exampleGraphQLApi", new Aws.AppSync.GraphQLApiArgs
        {
            LogConfig = new Aws.AppSync.Inputs.GraphQLApiLogConfigArgs
            {
                CloudwatchLogsRoleArn = exampleRole.Arn,
                FieldLogLevel = "ERROR",
            },
        });
    }

}

