using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var selected = Output.Create(Aws.Route53.GetZone.InvokeAsync(new Aws.Route53.GetZoneArgs
        {
            Name = "test.com.",
            PrivateZone = true,
        }));
        var www = new Aws.Route53.Record("www", new Aws.Route53.RecordArgs
        {
            Name = selected.Apply(selected => $"www.{selected.Name}"),
            Records = 
            {
                "10.0.0.1",
            },
            Ttl = 300,
            Type = "A",
            ZoneId = selected.Apply(selected => selected.ZoneId),
        });
    }

}

