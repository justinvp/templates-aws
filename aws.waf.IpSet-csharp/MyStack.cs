using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ipset = new Aws.Waf.IpSet("ipset", new Aws.Waf.IpSetArgs
        {
            IpSetDescriptors = 
            {
                new Aws.Waf.Inputs.IpSetIpSetDescriptorArgs
                {
                    Type = "IPV4",
                    Value = "192.0.7.0/24",
                },
                new Aws.Waf.Inputs.IpSetIpSetDescriptorArgs
                {
                    Type = "IPV4",
                    Value = "10.16.16.0/16",
                },
            },
        });
    }

}

