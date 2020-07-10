using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.Ec2.DefaultVpcDhcpOptions("default", new Aws.Ec2.DefaultVpcDhcpOptionsArgs
        {
            Tags = 
            {
                { "Name", "Default DHCP Option Set" },
            },
        });
    }

}

