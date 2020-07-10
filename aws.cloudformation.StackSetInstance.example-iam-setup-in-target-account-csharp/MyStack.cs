using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy = Output.Create(Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
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
                                aws_iam_role.AWSCloudFormationStackSetAdministrationRole.Arn,
                            },
                            Type = "AWS",
                        },
                    },
                },
            },
        }));
        var aWSCloudFormationStackSetExecutionRole = new Aws.Iam.Role("aWSCloudFormationStackSetExecutionRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy.Apply(aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy => aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy.Json),
        });
        var aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument = Output.Create(Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
        {
            Statements = 
            {
                new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
                {
                    Actions = 
                    {
                        "cloudformation:*",
                        "s3:*",
                        "sns:*",
                    },
                    Effect = "Allow",
                    Resources = 
                    {
                        "*",
                    },
                },
            },
        }));
        var aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyRolePolicy = new Aws.Iam.RolePolicy("aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyRolePolicy", new Aws.Iam.RolePolicyArgs
        {
            Policy = aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument.Apply(aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument => aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument.Json),
            Role = aWSCloudFormationStackSetExecutionRole.Name,
        });
    }

}

