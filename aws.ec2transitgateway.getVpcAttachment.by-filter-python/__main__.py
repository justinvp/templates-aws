import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_vpc_attachment(filters=[{
    "name": "vpc-id",
    "values": ["vpc-12345678"],
}])

