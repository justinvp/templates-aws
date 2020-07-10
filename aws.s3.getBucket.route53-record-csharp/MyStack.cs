using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var selected = Output.Create(Aws.S3.GetBucket.InvokeAsync(new Aws.S3.GetBucketArgs
        {
            Bucket = "bucket.test.com",
        }));
        var testZone = Output.Create(Aws.Route53.GetZone.InvokeAsync(new Aws.Route53.GetZoneArgs
        {
            Name = "test.com.",
        }));
        var example = new Aws.Route53.Record("example", new Aws.Route53.RecordArgs
        {
            Aliases = 
            {
                new Aws.Route53.Inputs.RecordAliasArgs
                {
                    Name = selected.Apply(selected => selected.WebsiteDomain),
                    ZoneId = selected.Apply(selected => selected.HostedZoneId),
                },
            },
            Name = "bucket",
            Type = "A",
            ZoneId = testZone.Apply(testZone => testZone.Id),
        });
    }

}

