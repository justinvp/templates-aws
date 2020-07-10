using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy = Output.Create(Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
        {
            Statements = 
            {
                new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
                {
                    Actions = 
                    {
                        "sts:AssumeRole",
                    },
                    Effect = "Allow",
                    Principals = 
                    {
                        new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
                        {
                            Identifiers = 
                            {
                                "cloudformation.amazonaws.com",
                            },
                            Type = "Service",
                        },
                    },
                },
            },
        }));
        var aWSCloudFormationStackSetAdministrationRole = new Aws.Iam.Role("aWSCloudFormationStackSetAdministrationRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy.Apply(aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy => aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy.Json),
        });
        var example = new Aws.CloudFormation.StackSet("example", new Aws.CloudFormation.StackSetArgs
        {
            AdministrationRoleArn = aWSCloudFormationStackSetAdministrationRole.Arn,
            Parameters = 
            {
                { "VPCCidr", "10.0.0.0/16" },
            },
            TemplateBody = @"{
  ""Parameters"" : {
    ""VPCCidr"" : {
      ""Type"" : ""String"",
      ""Default"" : ""10.0.0.0/16"",
      ""Description"" : ""Enter the CIDR block for the VPC. Default is 10.0.0.0/16.""
    }
  },
  ""Resources"" : {
    ""myVpc"": {
      ""Type"" : ""AWS::EC2::VPC"",
      ""Properties"" : {
        ""CidrBlock"" : { ""Ref"" : ""VPCCidr"" },
        ""Tags"" : [
          {""Key"": ""Name"", ""Value"": ""Primary_CF_VPC""}
        ]
      }
    }
  }
}

",
        });
        var aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument = example.ExecutionRoleName.Apply(executionRoleName => Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
        {
            Statements = 
            {
                new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
                {
                    Actions = 
                    {
                        "sts:AssumeRole",
                    },
                    Effect = "Allow",
                    Resources = 
                    {
                        $"arn:aws:iam::*:role/{executionRoleName}",
                    },
                },
            },
        }));
        var aWSCloudFormationStackSetAdministrationRoleExecutionPolicyRolePolicy = new Aws.Iam.RolePolicy("aWSCloudFormationStackSetAdministrationRoleExecutionPolicyRolePolicy", new Aws.Iam.RolePolicyArgs
        {
            Policy = aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument.Apply(aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument => aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument.Json),
            Role = aWSCloudFormationStackSetAdministrationRole.Name,
        });
    }

}

