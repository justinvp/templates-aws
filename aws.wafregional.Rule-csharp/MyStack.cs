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
        var wafrule = new Aws.WafRegional.Rule("wafrule", new Aws.WafRegional.RuleArgs
        {
            MetricName = "tfWAFRule",
            Predicates = 
            {
                new Aws.WafRegional.Inputs.RulePredicateArgs
                {
                    DataId = ipset.Id,
                    Negated = false,
                    Type = "IPMatch",
                },
            },
        });
    }

}

