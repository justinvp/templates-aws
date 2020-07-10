import pulumi
import pulumi_aws as aws

example_rule = aws.waf.Rule("exampleRule", metric_name="example")
example_rule_group = aws.waf.RuleGroup("exampleRuleGroup",
    activated_rules=[{
        "action": {
            "type": "COUNT",
        },
        "priority": 50,
        "rule_id": example_rule.id,
    }],
    metric_name="example")

