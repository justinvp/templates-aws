using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ipset = new Aws.WafRegional.IpSet("ipset", new Aws.WafRegional.IpSetArgs
        {
            IpSetDescriptors = 
            {
                new Aws.WafRegional.Inputs.IpSetIpSetDescriptorArgs
                {
                    Type = "IPV4",
                    Value = "192.0.7.0/24",
                },
            },
        });
        var wafrule = new Aws.WafRegional.RateBasedRule("wafrule", new Aws.WafRegional.RateBasedRuleArgs
        {
            MetricName = "tfWAFRule",
            Predicates = 
            {
                new Aws.WafRegional.Inputs.RateBasedRulePredicateArgs
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
                "aws_wafregional_ipset.ipset",
            },
        });
    }

}

