import pulumi
import pulumi_aws as aws

monitor = aws.opsworks.GangliaLayer("monitor",
    password="foobarbaz",
    stack_id=aws_opsworks_stack["main"]["id"])

