using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2.SecurityGroupRule("example", new Aws.Ec2.SecurityGroupRuleArgs
        {
            Type = "ingress",
            FromPort = 0,
            ToPort = 65535,
            Protocol = "tcp",
            CidrBlocks = aws_vpc.Example.Cidr_block,
            SecurityGroupId = "sg-123456",
        });
    }

}

