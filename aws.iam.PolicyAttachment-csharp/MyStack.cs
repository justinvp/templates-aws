using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var user = new Aws.Iam.User("user", new Aws.Iam.UserArgs
        {
        });
        var role = new Aws.Iam.Role("role", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Action"": ""sts:AssumeRole"",
      ""Principal"": {
        ""Service"": ""ec2.amazonaws.com""
      },
      ""Effect"": ""Allow"",
      ""Sid"": """"
    }
  ]
}

",
        });
        var @group = new Aws.Iam.Group("group", new Aws.Iam.GroupArgs
        {
        });
        var policy = new Aws.Iam.Policy("policy", new Aws.Iam.PolicyArgs
        {
            Description = "A test policy",
            Policy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Action"": [
        ""ec2:Describe*""
      ],
      ""Effect"": ""Allow"",
      ""Resource"": ""*""
    }
  ]
}

",
        });
        var test_attach = new Aws.Iam.PolicyAttachment("test-attach", new Aws.Iam.PolicyAttachmentArgs
        {
            Groups = 
            {
                @group.Name,
            },
            PolicyArn = policy.Arn,
            Roles = 
            {
                role.Name,
            },
            Users = 
            {
                user.Name,
            },
        });
    }

}

