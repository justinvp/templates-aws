import pulumi
import pulumi_aws as aws

example = aws.mq.Broker("example",
    broker_name="example",
    configuration={
        "id": aws_mq_configuration["test"]["id"],
        "revision": aws_mq_configuration["test"]["latest_revision"],
    },
    engine_type="ActiveMQ",
    engine_version="5.15.0",
    host_instance_type="mq.t2.micro",
    security_groups=[aws_security_group["test"]["id"]],
    users=[{
        "password": "MindTheGap",
        "username": "ExampleUser",
    }])

