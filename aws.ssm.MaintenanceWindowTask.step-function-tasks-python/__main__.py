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
    task_arn=aws_sfn_activity["example"]["id"],
    task_invocation_parameters={
        "stepFunctionsParameters": {
            "input": "{\"key1\":\"value1\"}",
            "name": "example",
        },
    },
    task_type="STEP_FUNCTIONS",
    window_id=aws_ssm_maintenance_window["example"]["id"])

