using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var key = new Aws.Kms.Key("key", new Aws.Kms.KeyArgs
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
        ""Service"": ""lambda.amazonaws.com""
      },
      ""Effect"": ""Allow"",
      ""Sid"": """"
    }
  ]
}

",
        });
        var grant = new Aws.Kms.Grant("grant", new Aws.Kms.GrantArgs
        {
            Constraints = 
            {
                new Aws.Kms.Inputs.GrantConstraintArgs
                {
                    EncryptionContextEquals = 
                    {
                        { "Department", "Finance" },
                    },
                },
            },
            GranteePrincipal = role.Arn,
            KeyId = key.KeyId,
            Operations = 
            {
                "Encrypt",
                "Decrypt",
                "GenerateDataKey",
            },
        });
    }

}

