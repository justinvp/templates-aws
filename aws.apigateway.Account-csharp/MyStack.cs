using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var cloudwatchRole = new Aws.Iam.Role("cloudwatchRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Sid"": """",
      ""Effect"": ""Allow"",
      ""Principal"": {
        ""Service"": ""apigateway.amazonaws.com""
      },
      ""Action"": ""sts:AssumeRole""
    }
  ]
}

",
        });
        var demo = new Aws.ApiGateway.Account("demo", new Aws.ApiGateway.AccountArgs
        {
            CloudwatchRoleArn = cloudwatchRole.Arn,
        });
        var cloudwatchRolePolicy = new Aws.Iam.RolePolicy("cloudwatchRolePolicy", new Aws.Iam.RolePolicyArgs
        {
            Policy = @"{
    ""Version"": ""2012-10-17"",
    ""Statement"": [
        {
            ""Effect"": ""Allow"",
            ""Action"": [
                ""logs:CreateLogGroup"",
                ""logs:CreateLogStream"",
                ""logs:DescribeLogGroups"",
                ""logs:DescribeLogStreams"",
                ""logs:PutLogEvents"",
                ""logs:GetLogEvents"",
                ""logs:FilterLogEvents""
            ],
            ""Resource"": ""*""
        }
    ]
}

",
            Role = cloudwatchRole.Id,
        });
    }

}

