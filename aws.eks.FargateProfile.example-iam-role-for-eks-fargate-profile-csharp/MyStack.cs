using System.Collections.Generic;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Iam.Role("example", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                { "Statement", new[]
                    {
                        new Dictionary<string, object?>
                        {
                            { "Action", "sts:AssumeRole" },
                            { "Effect", "Allow" },
                            { "Principal", new Dictionary<string, object?>
                            {
                                { "Service", "eks-fargate-pods.amazonaws.com" },
                            } },
                        },
                    }
                 },
                { "Version", "2012-10-17" },
            }),
        });
        var example_AmazonEKSFargatePodExecutionRolePolicy = new Aws.Iam.RolePolicyAttachment("example-AmazonEKSFargatePodExecutionRolePolicy", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = "arn:aws:iam::aws:policy/AmazonEKSFargatePodExecutionRolePolicy",
            Role = example.Name,
        });
    }

}

