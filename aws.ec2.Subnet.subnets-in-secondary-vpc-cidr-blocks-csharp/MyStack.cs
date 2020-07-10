using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var secondaryCidr = new Aws.Ec2.VpcIpv4CidrBlockAssociation("secondaryCidr", new Aws.Ec2.VpcIpv4CidrBlockAssociationArgs
        {
            CidrBlock = "172.2.0.0/16",
            VpcId = aws_vpc.Main.Id,
        });
        var inSecondaryCidr = new Aws.Ec2.Subnet("inSecondaryCidr", new Aws.Ec2.SubnetArgs
        {
            CidrBlock = "172.2.0.0/24",
            VpcId = secondaryCidr.VpcId,
        });
    }

}

