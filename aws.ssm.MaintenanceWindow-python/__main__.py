import pulumi
import pulumi_aws as aws

production = aws.ssm.MaintenanceWindow("production",
    cutoff=1,
    duration=3,
    schedule="cron(0 16 ? * TUE *)")

