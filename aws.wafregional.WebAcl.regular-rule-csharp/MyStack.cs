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
        var wafacl = new Aws.WafRegional.WebAcl("wafacl", new Aws.WafRegional.WebAclArgs
        {
            DefaultAction = new Aws.WafRegional.Inputs.WebAclDefaultActionArgs
            {
                Type = "ALLOW",
            },
            MetricName = "tfWebACL",
            Rules = 
            {
                new Aws.WafRegional.Inputs.WebAclRuleArgs
                {
                    Action = new Aws.WafRegional.Inputs.WebAclRuleActionArgs
                    {
                        Type = "BLOCK",
                    },
                    Priority = 1,
                    RuleId = wafrule.Id,
                    Type = "REGULAR",
                },
            },
        });
    }

}

