using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleZone = new Aws.Route53.Zone("exampleZone", new Aws.Route53.ZoneArgs
        {
        });
        var exampleRecord = new Aws.Route53.Record("exampleRecord", new Aws.Route53.RecordArgs
        {
            AllowOverwrite = true,
            Name = "test.example.com",
            Records = 
            {
                exampleZone.NameServers.Apply(nameServers => nameServers[0]),
                exampleZone.NameServers.Apply(nameServers => nameServers[1]),
                exampleZone.NameServers.Apply(nameServers => nameServers[2]),
                exampleZone.NameServers.Apply(nameServers => nameServers[3]),
            },
            Ttl = 30,
            Type = "NS",
            ZoneId = exampleZone.ZoneId,
        });
    }

}

