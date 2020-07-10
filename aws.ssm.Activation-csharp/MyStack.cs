using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testRole = new Aws.Iam.Role("testRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"  {
    ""Version"": ""2012-10-17"",
    ""Statement"": {
      ""Effect"": ""Allow"",
      ""Principal"": {""Service"": ""ssm.amazonaws.com""},
      ""Action"": ""sts:AssumeRole""
    }
  }

",
        });
        var testAttach = new Aws.Iam.RolePolicyAttachment("testAttach", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore",
            Role = testRole.Name,
        });
        var foo = new Aws.Ssm.Activation("foo", new Aws.Ssm.ActivationArgs
        {
            Description = "Test",
            IamRole = testRole.Id,
            RegistrationLimit = 5,
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_iam_role_policy_attachment.test_attach",
            },
        });
    }

}

