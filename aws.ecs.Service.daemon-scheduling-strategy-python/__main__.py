import pulumi
import pulumi_aws as aws

bar = aws.ecs.Service("bar",
    cluster=aws_ecs_cluster["foo"]["id"],
    scheduling_strategy="DAEMON",
    task_definition=aws_ecs_task_definition["bar"]["arn"])

