using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.WafRegional.WebAcl("example", new Aws.WafRegional.WebAclArgs
        {
            DefaultAction = new Aws.WafRegional.Inputs.WebAclDefaultActionArgs
            {
                Type = "ALLOW",
            },
            MetricName = "example",
            Rules = 
            {
                new Aws.WafRegional.Inputs.WebAclRuleArgs
                {
                    OverrideAction = new Aws.WafRegional.Inputs.WebAclRuleOverrideActionArgs
                    {
                        Type = "NONE",
                    },
                    Priority = 1,
                    RuleId = aws_wafregional_rule_group.Example.Id,
                    Type = "GROUP",
                },
            },
        });
    }

}

