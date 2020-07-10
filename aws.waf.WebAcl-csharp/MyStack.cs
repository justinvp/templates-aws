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
        var wafAcl = new Aws.Waf.WebAcl("wafAcl", new Aws.Waf.WebAclArgs
        {
            DefaultAction = new Aws.Waf.Inputs.WebAclDefaultActionArgs
            {
                Type = "ALLOW",
            },
            MetricName = "tfWebACL",
            Rules = 
            {
                new Aws.Waf.Inputs.WebAclRuleArgs
                {
                    Action = new Aws.Waf.Inputs.WebAclRuleActionArgs
                    {
                        Type = "BLOCK",
                    },
                    Priority = 1,
                    RuleId = wafrule.Id,
                    Type = "REGULAR",
                },
            },
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_waf_ipset.ipset",
                "aws_waf_rule.wafrule",
            },
        });
    }

}

