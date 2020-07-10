using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var snEc2 = new Aws.Ec2.VpcEndpointSubnetAssociation("snEc2", new Aws.Ec2.VpcEndpointSubnetAssociationArgs
        {
            SubnetId = aws_subnet.Sn.Id,
            VpcEndpointId = aws_vpc_endpoint.Ec2.Id,
        });
    }

}

