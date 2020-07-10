using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var dnsResolver = new Aws.Ec2.VpcDhcpOptions("dnsResolver", new Aws.Ec2.VpcDhcpOptionsArgs
        {
            DomainNameServers = 
            {
                "8.8.8.8",
                "8.8.4.4",
            },
        });
    }

}

