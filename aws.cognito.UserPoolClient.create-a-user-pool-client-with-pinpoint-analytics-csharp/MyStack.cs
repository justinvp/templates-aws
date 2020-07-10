using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var current = Output.Create(Aws.GetCallerIdentity.InvokeAsync());
        var testUserPool = new Aws.Cognito.UserPool("testUserPool", new Aws.Cognito.UserPoolArgs
        {
        });
        var testApp = new Aws.Pinpoint.App("testApp", new Aws.Pinpoint.AppArgs
        {
        });
        var testRole = new Aws.Iam.Role("testRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Action"": ""sts:AssumeRole"",
      ""Principal"": {
        ""Service"": ""cognito-idp.amazonaws.com""
      },
      ""Effect"": ""Allow"",
      ""Sid"": """"
    }
  ]
}

",
        });
        var testRolePolicy = new Aws.Iam.RolePolicy("testRolePolicy", new Aws.Iam.RolePolicyArgs
        {
            Policy = Output.Tuple(current, testApp.ApplicationId).Apply(values =>
            {
                var current = values.Item1;
                var applicationId = values.Item2;
                return @$"{{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {{
      ""Action"": [
        ""mobiletargeting:UpdateEndpoint"",
        ""mobiletargeting:PutItems""
      ],
      ""Effect"": ""Allow"",
      ""Resource"": ""arn:aws:mobiletargeting:*:{current.AccountId}:apps/{applicationId}*""
    }}
  ]
}}

";
            }),
            Role = testRole.Id,
        });
        var testUserPoolClient = new Aws.Cognito.UserPoolClient("testUserPoolClient", new Aws.Cognito.UserPoolClientArgs
        {
            AnalyticsConfiguration = new Aws.Cognito.Inputs.UserPoolClientAnalyticsConfigurationArgs
            {
                ApplicationId = testApp.ApplicationId,
                ExternalId = "some_id",
                RoleArn = testRole.Arn,
                UserDataShared = true,
            },
            UserPoolId = testUserPool.Id,
        });
    }

}

