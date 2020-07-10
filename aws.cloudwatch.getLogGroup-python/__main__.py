import pulumi
import pulumi_aws as aws

example = aws.cloudwatch.get_log_group(name="MyImportantLogs")

