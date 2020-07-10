import pulumi
import pulumi_aws as aws

dev_account_access = aws.cloudwatch.EventPermission("devAccountAccess",
    principal="123456789012",
    statement_id="DevAccountAccess")

