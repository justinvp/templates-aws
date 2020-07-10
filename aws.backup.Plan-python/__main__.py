import pulumi
import pulumi_aws as aws

example = aws.backup.Plan("example", rules=[{
    "rule_name": "tf_example_backup_rule",
    "schedule": "cron(0 12 * * ? *)",
    "targetVaultName": aws_backup_vault["test"]["name"],
}])

