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
        var wafrule = new Aws.Waf.Rule("wafrule", new Aws.Waf.RuleArgs
        {
            MetricName = "tfWAFRule",
            Predicates = 
            {
                new Aws.Waf.Inputs.RulePredicateArgs
                {
                    DataId = ipset.Id,
                    Negated = false,
                    Type = "IPMatch",
                },
            },
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_waf_ipset.ipset",
            },
        });
    }

}

