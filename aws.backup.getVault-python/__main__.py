import pulumi
import pulumi_aws as aws

example = aws.backup.get_vault(name="example_backup_vault")

