using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleRule = new Aws.Waf.Rule("exampleRule", new Aws.Waf.RuleArgs
        {
            MetricName = "example",
        });
        var exampleRuleGroup = new Aws.Waf.RuleGroup("exampleRuleGroup", new Aws.Waf.RuleGroupArgs
        {
            ActivatedRules = 
            {
                new Aws.Waf.Inputs.RuleGroupActivatedRuleArgs
                {
                    Action = new Aws.Waf.Inputs.RuleGroupActivatedRuleActionArgs
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

