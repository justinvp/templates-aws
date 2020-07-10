import pulumi
import pulumi_aws as aws

example = aws.ssm.MaintenanceWindowTask("example",
    max_concurrency=2,
    max_errors=1,
    priority=1,
    service_role_arn=aws_iam_role["example"]["arn"],
    targets=[{
        "key": "InstanceIds",
        "values": [aws_instance["example"]["id"]],
    }],
    task_arn="AWS-RunShellScript",
    task_invocation_parameters={
        "runCommandParameters": {
            "notificationConfig": {
                "notificationArn": aws_sns_topic["example"]["arn"],
                "notificationEvents": ["All"],
                "notification_type": "Command",
            },
            "outputS3Bucket": aws_s3_bucket["example"]["bucket"],
            "outputS3KeyPrefix": "output",
            "parameter": [{
                "name": "commands",
                "values": ["date"],
            }],
            "service_role_arn": aws_iam_role["example"]["arn"],
            "timeoutSeconds": 600,
        },
    },
    task_type="RUN_COMMAND",
    window_id=aws_ssm_maintenance_window["example"]["id"])

