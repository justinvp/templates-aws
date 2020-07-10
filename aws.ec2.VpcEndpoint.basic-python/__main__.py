import pulumi
import pulumi_aws as aws

s3 = aws.ec2.VpcEndpoint("s3",
    service_name="com.amazonaws.us-west-2.s3",
    vpc_id=aws_vpc["main"]["id"])

