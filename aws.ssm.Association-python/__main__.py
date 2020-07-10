import pulumi
import pulumi_aws as aws

example = aws.ssm.Association("example", targets=[{
    "key": "InstanceIds",
    "values": [aws_instance["example"]["id"]],
}])

