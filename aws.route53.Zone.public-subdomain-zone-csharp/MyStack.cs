using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var main = new Aws.Route53.Zone("main", new Aws.Route53.ZoneArgs
        {
        });
        var dev = new Aws.Route53.Zone("dev", new Aws.Route53.ZoneArgs
        {
            Tags = 
            {
                { "Environment", "dev" },
            },
        });
        var dev_ns = new Aws.Route53.Record("dev-ns", new Aws.Route53.RecordArgs
        {
            Name = "dev.example.com",
            Records = 
            {
                dev.NameServers.Apply(nameServers => nameServers[0]),
                dev.NameServers.Apply(nameServers => nameServers[1]),
                dev.NameServers.Apply(nameServers => nameServers[2]),
                dev.NameServers.Apply(nameServers => nameServers[3]),
            },
            Ttl = 30,
            Type = "NS",
            ZoneId = main.ZoneId,
        });
    }

}

