using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var www_dev = new Aws.Route53.Record("www-dev", new Aws.Route53.RecordArgs
        {
            Name = "www",
            Records = 
            {
                "dev.example.com",
            },
            SetIdentifier = "dev",
            Ttl = 5,
            Type = "CNAME",
            WeightedRoutingPolicies = 
            {
                new Aws.Route53.Inputs.RecordWeightedRoutingPolicyArgs
                {
                    Weight = 10,
                },
            },
            ZoneId = aws_route53_zone.Primary.Zone_id,
        });
        var www_live = new Aws.Route53.Record("www-live", new Aws.Route53.RecordArgs
        {
            Name = "www",
            Records = 
            {
                "live.example.com",
            },
            SetIdentifier = "live",
            Ttl = 5,
            Type = "CNAME",
            WeightedRoutingPolicies = 
            {
                new Aws.Route53.Inputs.RecordWeightedRoutingPolicyArgs
                {
                    Weight = 90,
                },
            },
            ZoneId = aws_route53_zone.Primary.Zone_id,
        });
    }

}

