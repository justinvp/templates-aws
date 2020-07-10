using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2ClientVpn.NetworkAssociation("example", new Aws.Ec2ClientVpn.NetworkAssociationArgs
        {
            ClientVpnEndpointId = aws_ec2_client_vpn_endpoint.Example.Id,
            SubnetId = aws_subnet.Example.Id,
        });
    }

}

