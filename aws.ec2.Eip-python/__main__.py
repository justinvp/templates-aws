import pulumi
import pulumi_aws as aws

lb = aws.ec2.Eip("lb",
    instance=aws_instance["web"]["id"],
    vpc=True)

