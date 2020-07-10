using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var mainIdentityPool = new Aws.Cognito.IdentityPool("mainIdentityPool", new Aws.Cognito.IdentityPoolArgs
        {
            AllowUnauthenticatedIdentities = false,
            IdentityPoolName = "identity pool",
            SupportedLoginProviders = 
            {
                { "graph.facebook.com", "7346241598935555" },
            },
        });
        var authenticatedRole = new Aws.Iam.Role("authenticatedRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = mainIdentityPool.Id.Apply(id => @$"{{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {{
      ""Effect"": ""Allow"",
      ""Principal"": {{
        ""Federated"": ""cognito-identity.amazonaws.com""
      }},
      ""Action"": ""sts:AssumeRoleWithWebIdentity"",
      ""Condition"": {{
        ""StringEquals"": {{
          ""cognito-identity.amazonaws.com:aud"": ""{id}""
        }},
        ""ForAnyValue:StringLike"": {{
          ""cognito-identity.amazonaws.com:amr"": ""authenticated""
        }}
      }}
    }}
  ]
}}

"),
        });
        var authenticatedRolePolicy = new Aws.Iam.RolePolicy("authenticatedRolePolicy", new Aws.Iam.RolePolicyArgs
        {
            Policy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Effect"": ""Allow"",
      ""Action"": [
        ""mobileanalytics:PutEvents"",
        ""cognito-sync:*"",
        ""cognito-identity:*""
      ],
      ""Resource"": [
        ""*""
      ]
    }
  ]
}

",
            Role = authenticatedRole.Id,
        });
        var mainIdentityPoolRoleAttachment = new Aws.Cognito.IdentityPoolRoleAttachment("mainIdentityPoolRoleAttachment", new Aws.Cognito.IdentityPoolRoleAttachmentArgs
        {
            IdentityPoolId = mainIdentityPool.Id,
            RoleMappings = 
            {
                new Aws.Cognito.Inputs.IdentityPoolRoleAttachmentRoleMappingArgs
                {
                    AmbiguousRoleResolution = "AuthenticatedRole",
                    IdentityProvider = "graph.facebook.com",
                    MappingRule = 
                    {
                        
                        {
                            { "claim", "isAdmin" },
                            { "matchType", "Equals" },
                            { "roleArn", authenticatedRole.Arn },
                            { "value", "paid" },
                        },
                    },
                    Type = "Rules",
                },
            },
            Roles = 
            {
                { "authenticated", authenticatedRole.Arn },
            },
        });
    }

}

