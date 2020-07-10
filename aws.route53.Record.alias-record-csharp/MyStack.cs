using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var main = new Aws.Elb.LoadBalancer("main", new Aws.Elb.LoadBalancerArgs
        {
            AvailabilityZones = 
            {
                "us-east-1c",
            },
            Listeners = 
            {
                new Aws.Elb.Inputs.LoadBalancerListenerArgs
                {
                    InstancePort = 80,
                    InstanceProtocol = "http",
                    LbPort = 80,
                    LbProtocol = "http",
                },
            },
        });
        var www = new Aws.Route53.Record("www", new Aws.Route53.RecordArgs
        {
            Aliases = 
            {
                new Aws.Route53.Inputs.RecordAliasArgs
                {
                    EvaluateTargetHealth = true,
                    Name = main.DnsName,
                    ZoneId = main.ZoneId,
                },
            },
            Name = "example.com",
            Type = "A",
            ZoneId = aws_route53_zone.Primary.Zone_id,
        });
    }

}

