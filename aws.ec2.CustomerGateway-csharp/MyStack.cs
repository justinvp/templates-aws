using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var main = new Aws.Ec2.CustomerGateway("main", new Aws.Ec2.CustomerGatewayArgs
        {
            BgpAsn = 65000,
            IpAddress = "172.83.124.10",
            Tags = 
            {
                { "Name", "main-customer-gateway" },
            },
            Type = "ipsec.1",
        });
    }

}

