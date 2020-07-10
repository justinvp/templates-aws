using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ipset = new Aws.Waf.IpSet("ipset", new Aws.Waf.IpSetArgs
        {
            IpSetDescriptors = 
            {
                new Aws.Waf.Inputs.IpSetIpSetDescriptorArgs
                {
                    Type = "IPV4",
                    Value = "192.0.7.0/24",
                },
            },
        });
        var wafrule = new Aws.Waf.RateBasedRule("wafrule", new Aws.Waf.RateBasedRuleArgs
        {
            MetricName = "tfWAFRule",
            Predicates = 
            {
                new Aws.Waf.Inputs.RateBasedRulePredicateArgs
                {
                    DataId = ipset.Id,
                    Negated = false,
                    Type = "IPMatch",
                },
            },
            RateKey = "IP",
            RateLimit = 100,
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_waf_ipset.ipset",
            },
        });
    }

}

