using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.WafV2.IpSet("example", new Aws.WafV2.IpSetArgs
        {
            Addresses = 
            {
                "1.2.3.4/32",
                "5.6.7.8/32",
            },
            Description = "Example IP set",
            IpAddressVersion = "IPV4",
            Scope = "REGIONAL",
            Tags = 
            {
                { "Tag1", "Value1" },
                { "Tag2", "Value2" },
            },
        });
    }

}

