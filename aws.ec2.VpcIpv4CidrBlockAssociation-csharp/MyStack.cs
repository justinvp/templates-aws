using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var main = new Aws.Ec2.Vpc("main", new Aws.Ec2.VpcArgs
        {
            CidrBlock = "10.0.0.0/16",
        });
        var secondaryCidr = new Aws.Ec2.VpcIpv4CidrBlockAssociation("secondaryCidr", new Aws.Ec2.VpcIpv4CidrBlockAssociationArgs
        {
            CidrBlock = "172.2.0.0/16",
            VpcId = main.Id,
        });
    }

}

