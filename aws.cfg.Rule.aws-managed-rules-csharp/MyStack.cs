using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var rule = new Aws.Cfg.Rule("rule", new Aws.Cfg.RuleArgs
        {
            Source = new Aws.Cfg.Inputs.RuleSourceArgs
            {
                Owner = "AWS",
                SourceIdentifier = "S3_BUCKET_VERSIONING_ENABLED",
            },
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_config_configuration_recorder.foo",
            },
        });
        var role = new Aws.Iam.Role("role", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Action"": ""sts:AssumeRole"",
      ""Principal"": {
        ""Service"": ""config.amazonaws.com""
      },
      ""Effect"": ""Allow"",
      ""Sid"": """"
    }
  ]
}

",
        });
        var foo = new Aws.Cfg.Recorder("foo", new Aws.Cfg.RecorderArgs
        {
            RoleArn = role.Arn,
        });
        var rolePolicy = new Aws.Iam.RolePolicy("rolePolicy", new Aws.Iam.RolePolicyArgs
        {
            Policy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
  	{
  		""Action"": ""config:Put*"",
  		""Effect"": ""Allow"",
  		""Resource"": ""*""

  	}
  ]
}

",
            Role = role.Id,
        });
    }

}

