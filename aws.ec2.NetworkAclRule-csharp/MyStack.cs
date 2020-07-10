using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var barNetworkAcl = new Aws.Ec2.NetworkAcl("barNetworkAcl", new Aws.Ec2.NetworkAclArgs
        {
            VpcId = aws_vpc.Foo.Id,
        });
        var barNetworkAclRule = new Aws.Ec2.NetworkAclRule("barNetworkAclRule", new Aws.Ec2.NetworkAclRuleArgs
        {
            NetworkAclId = barNetworkAcl.Id,
            RuleNumber = 200,
            Egress = false,
            Protocol = "tcp",
            RuleAction = "allow",
            CidrBlock = aws_vpc.Foo.Cidr_block,
            FromPort = 22,
            ToPort = 22,
        });
    }

}

