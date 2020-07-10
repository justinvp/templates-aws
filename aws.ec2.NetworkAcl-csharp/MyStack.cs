using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var main = new Aws.Ec2.NetworkAcl("main", new Aws.Ec2.NetworkAclArgs
        {
            Egress = 
            {
                new Aws.Ec2.Inputs.NetworkAclEgressArgs
                {
                    Action = "allow",
                    CidrBlock = "10.3.0.0/18",
                    FromPort = 443,
                    Protocol = "tcp",
                    RuleNo = 200,
                    ToPort = 443,
                },
            },
            Ingress = 
            {
                new Aws.Ec2.Inputs.NetworkAclIngressArgs
                {
                    Action = "allow",
                    CidrBlock = "10.3.0.0/18",
                    FromPort = 80,
                    Protocol = "tcp",
                    RuleNo = 100,
                    ToPort = 80,
                },
            },
            Tags = 
            {
                { "Name", "main" },
            },
            VpcId = aws_vpc.Main.Id,
        });
    }

}

