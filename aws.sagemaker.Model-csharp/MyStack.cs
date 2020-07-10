using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var model = new Aws.Sagemaker.Model("model", new Aws.Sagemaker.ModelArgs
        {
            ExecutionRoleArn = aws_iam_role.Foo.Arn,
            PrimaryContainer = new Aws.Sagemaker.Inputs.ModelPrimaryContainerArgs
            {
                Image = "174872318107.dkr.ecr.us-west-2.amazonaws.com/kmeans:1",
            },
        });
        var assumeRole = Output.Create(Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
        {
            Statements = 
            {
                new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
                {
                    Actions = 
                    {
                        "sts:AssumeRole",
                    },
                    Principals = 
                    {
                        new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
                        {
                            Identifiers = 
                            {
                                "sagemaker.amazonaws.com",
                            },
                            Type = "Service",
                        },
                    },
                },
            },
        }));
        var role = new Aws.Iam.Role("role", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = assumeRole.Apply(assumeRole => assumeRole.Json),
        });
    }

}

