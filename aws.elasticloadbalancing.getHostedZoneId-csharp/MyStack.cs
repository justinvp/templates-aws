using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var main = Output.Create(Aws.Elb.GetHostedZoneId.InvokeAsync());
        var www = new Aws.Route53.Record("www", new Aws.Route53.RecordArgs
        {
            Aliases = 
            {
                new Aws.Route53.Inputs.RecordAliasArgs
                {
                    EvaluateTargetHealth = true,
                    Name = aws_elb.Main.Dns_name,
                    ZoneId = main.Apply(main => main.Id),
                },
            },
            Name = "example.com",
            Type = "A",
            ZoneId = aws_route53_zone.Primary.Zone_id,
        });
    }

}

