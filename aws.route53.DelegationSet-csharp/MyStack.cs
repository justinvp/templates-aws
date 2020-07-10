using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var main = new Aws.Route53.DelegationSet("main", new Aws.Route53.DelegationSetArgs
        {
            ReferenceName = "DynDNS",
        });
        var primary = new Aws.Route53.Zone("primary", new Aws.Route53.ZoneArgs
        {
            DelegationSetId = main.Id,
        });
        var secondary = new Aws.Route53.Zone("secondary", new Aws.Route53.ZoneArgs
        {
            DelegationSetId = main.Id,
        });
    }

}

