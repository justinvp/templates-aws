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
                                { "Service", "ec2.amazonaws.com" },
                            } },
                        },
                    }
                 },
                { "Version", "2012-10-17" },
            }),
        });
        var example_AmazonEKSWorkerNodePolicy = new Aws.Iam.RolePolicyAttachment("example-AmazonEKSWorkerNodePolicy", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
            Role = example.Name,
        });
        var example_AmazonEKSCNIPolicy = new Aws.Iam.RolePolicyAttachment("example-AmazonEKSCNIPolicy", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy",
            Role = example.Name,
        });
        var example_AmazonEC2ContainerRegistryReadOnly = new Aws.Iam.RolePolicyAttachment("example-AmazonEC2ContainerRegistryReadOnly", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
            Role = example.Name,
        });
    }

}

