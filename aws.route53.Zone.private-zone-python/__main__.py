import pulumi
import pulumi_aws as aws

private = aws.route53.Zone("private", vpcs=[{
    "vpc_id": aws_vpc["example"]["id"],
}])

