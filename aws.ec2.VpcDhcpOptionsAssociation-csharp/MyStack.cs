using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var dnsResolver = new Aws.Ec2.VpcDhcpOptionsAssociation("dnsResolver", new Aws.Ec2.VpcDhcpOptionsAssociationArgs
        {
            DhcpOptionsId = aws_vpc_dhcp_options.Foo.Id,
            VpcId = aws_vpc.Foo.Id,
        });
    }

}

