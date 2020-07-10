using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var role = new Aws.Iam.Role("role", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"    {
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
        var test_attach = new Aws.Iam.RolePolicyAttachment("test-attach", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = policy.Arn,
            Role = role.Name,
        });
    }

}

