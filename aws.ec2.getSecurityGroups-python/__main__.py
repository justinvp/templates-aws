import pulumi
import pulumi_aws as aws

test = aws.ec2.get_security_groups(tags={
    "Application": "k8s",
    "Environment": "dev",
})

