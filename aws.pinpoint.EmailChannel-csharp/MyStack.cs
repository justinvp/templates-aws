using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var app = new Aws.Pinpoint.App("app", new Aws.Pinpoint.AppArgs
        {
        });
        var identity = new Aws.Ses.DomainIdentity("identity", new Aws.Ses.DomainIdentityArgs
        {
            Domain = "example.com",
        });
        var role = new Aws.Iam.Role("role", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Action"": ""sts:AssumeRole"",
      ""Principal"": {
        ""Service"": ""pinpoint.amazonaws.com""
      },
      ""Effect"": ""Allow"",
      ""Sid"": """"
    }
  ]
}

",
        });
        var email = new Aws.Pinpoint.EmailChannel("email", new Aws.Pinpoint.EmailChannelArgs
        {
            ApplicationId = app.ApplicationId,
            FromAddress = "user@example.com",
            Identity = identity.Arn,
            RoleArn = role.Arn,
        });
        var rolePolicy = new Aws.Iam.RolePolicy("rolePolicy", new Aws.Iam.RolePolicyArgs
        {
            Policy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": {
    ""Action"": [
      ""mobileanalytics:PutEvents"",
      ""mobileanalytics:PutItems""
    ],
    ""Effect"": ""Allow"",
    ""Resource"": [
      ""*""
    ]
  }
}

",
            Role = role.Id,
        });
    }

}

