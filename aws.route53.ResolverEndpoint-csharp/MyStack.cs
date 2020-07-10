using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Aws.Route53.ResolverEndpoint("foo", new Aws.Route53.ResolverEndpointArgs
        {
            Direction = "INBOUND",
            IpAddresses = 
            {
                new Aws.Route53.Inputs.ResolverEndpointIpAddressArgs
                {
                    SubnetId = aws_subnet.Sn1.Id,
                },
                new Aws.Route53.Inputs.ResolverEndpointIpAddressArgs
                {
                    Ip = "10.0.64.4",
                    SubnetId = aws_subnet.Sn2.Id,
                },
            },
            SecurityGroupIds = 
            {
                aws_security_group.Sg1.Id,
                aws_security_group.Sg2.Id,
            },
            Tags = 
            {
                { "Environment", "Prod" },
            },
        });
    }

}

