import pulumi
import pulumi_aws as aws

example = aws.backup.get_plan(plan_id="tf_example_backup_plan_id")

