using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var mainUserPool = new Aws.Cognito.UserPool("mainUserPool", new Aws.Cognito.UserPoolArgs
        {
        });
        var groupRole = new Aws.Iam.Role("groupRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Sid"": """",
      ""Effect"": ""Allow"",
      ""Principal"": {
        ""Federated"": ""cognito-identity.amazonaws.com""
      },
      ""Action"": ""sts:AssumeRoleWithWebIdentity"",
      ""Condition"": {
        ""StringEquals"": {
          ""cognito-identity.amazonaws.com:aud"": ""us-east-1:12345678-dead-beef-cafe-123456790ab""
        },
        ""ForAnyValue:StringLike"": {
          ""cognito-identity.amazonaws.com:amr"": ""authenticated""
        }
      }
    }
  ]
}

",
        });
        var mainUserGroup = new Aws.Cognito.UserGroup("mainUserGroup", new Aws.Cognito.UserGroupArgs
        {
            Description = "Managed by Pulumi",
            Precedence = 42,
            RoleArn = groupRole.Arn,
            UserPoolId = mainUserPool.Id,
        });
    }

}

