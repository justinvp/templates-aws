import pulumi
import pulumi_aws as aws

ecs_mongo = aws.ecs.get_container_definition(container_name="mongodb",
    task_definition=aws_ecs_task_definition["mongo"]["id"])

