using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var organizationRole = new Aws.Iam.Role("organizationRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Sid"": """",
      ""Effect"": ""Allow"",
      ""Principal"": {
        ""Service"": ""config.amazonaws.com""
      },
      ""Action"": ""sts:AssumeRole""
    }
  ]
}

",
        });
        var organizationConfigurationAggregator = new Aws.Cfg.ConfigurationAggregator("organizationConfigurationAggregator", new Aws.Cfg.ConfigurationAggregatorArgs
        {
            OrganizationAggregationSource = new Aws.Cfg.Inputs.ConfigurationAggregatorOrganizationAggregationSourceArgs
            {
                AllRegions = true,
                RoleArn = organizationRole.Arn,
            },
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_iam_role_policy_attachment.organization",
            },
        });
        var organizationRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("organizationRolePolicyAttachment", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = "arn:aws:iam::aws:policy/service-role/AWSConfigRoleForOrganizations",
            Role = organizationRole.Name,
        });
    }

}

