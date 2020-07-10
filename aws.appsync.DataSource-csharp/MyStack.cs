using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleTable = new Aws.DynamoDB.Table("exampleTable", new Aws.DynamoDB.TableArgs
        {
            Attributes = 
            {
                new Aws.DynamoDB.Inputs.TableAttributeArgs
                {
                    Name = "UserId",
                    Type = "S",
                },
            },
            HashKey = "UserId",
            ReadCapacity = 1,
            WriteCapacity = 1,
        });
        var exampleRole = new Aws.Iam.Role("exampleRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Action"": ""sts:AssumeRole"",
      ""Principal"": {
        ""Service"": ""appsync.amazonaws.com""
      },
      ""Effect"": ""Allow""
    }
  ]
}

",
        });
        var exampleRolePolicy = new Aws.Iam.RolePolicy("exampleRolePolicy", new Aws.Iam.RolePolicyArgs
        {
            Policy = exampleTable.Arn.Apply(arn => @$"{{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {{
      ""Action"": [
        ""dynamodb:*""
      ],
      ""Effect"": ""Allow"",
      ""Resource"": [
        ""{arn}""
      ]
    }}
  ]
}}

"),
            Role = exampleRole.Id,
        });
        var exampleGraphQLApi = new Aws.AppSync.GraphQLApi("exampleGraphQLApi", new Aws.AppSync.GraphQLApiArgs
        {
            AuthenticationType = "API_KEY",
        });
        var exampleDataSource = new Aws.AppSync.DataSource("exampleDataSource", new Aws.AppSync.DataSourceArgs
        {
            ApiId = exampleGraphQLApi.Id,
            DynamodbConfig = new Aws.AppSync.Inputs.DataSourceDynamodbConfigArgs
            {
                TableName = exampleTable.Name,
            },
            ServiceRoleArn = exampleRole.Arn,
            Type = "AMAZON_DYNAMODB",
        });
    }

}

