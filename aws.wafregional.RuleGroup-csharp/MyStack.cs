using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleRule = new Aws.WafRegional.Rule("exampleRule", new Aws.WafRegional.RuleArgs
        {
            MetricName = "example",
        });
        var exampleRuleGroup = new Aws.WafRegional.RuleGroup("exampleRuleGroup", new Aws.WafRegional.RuleGroupArgs
        {
            ActivatedRules = 
            {
                new Aws.WafRegional.Inputs.RuleGroupActivatedRuleArgs
                {
                    Action = new Aws.WafRegional.Inputs.RuleGroupActivatedRuleActionArgs
                    {
                        Type = "COUNT",
                    },
                    Priority = 50,
                    RuleId = exampleRule.Id,
                },
            },
            MetricName = "example",
        });
    }

}

