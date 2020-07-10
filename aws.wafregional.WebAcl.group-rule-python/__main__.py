import pulumi
import pulumi_aws as aws

example = aws.wafregional.WebAcl("example",
    default_action={
        "type": "ALLOW",
    },
    metric_name="example",
    rules=[{
        "overrideAction": {
            "type": "NONE",
        },
        "priority": 1,
        "rule_id": aws_wafregional_rule_group["example"]["id"],
        "type": "GROUP",
    }])

