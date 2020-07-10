using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var primary = new Aws.Route53.Zone("primary", new Aws.Route53.ZoneArgs
        {
        });
    }

}

