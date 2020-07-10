import pulumi
import pulumi_aws as aws

main = aws.ses.ActiveReceiptRuleSet("main", rule_set_name="primary-rules")

