using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.LicenseManager.LicenseConfiguration("example", new Aws.LicenseManager.LicenseConfigurationArgs
        {
            Description = "Example",
            LicenseCount = 10,
            LicenseCountHardLimit = true,
            LicenseCountingType = "Socket",
            LicenseRules = 
            {
                "#minimumSockets=2",
            },
            Tags = 
            {
                { "foo", "barr" },
            },
        });
    }

}

