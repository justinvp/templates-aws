import pulumi
import pulumi_aws as aws

s3 = aws.ec2.get_vpc_endpoint_service(service="s3")
# Create a VPC
foo = aws.ec2.Vpc("foo", cidr_block="10.0.0.0/16")
# Create a VPC endpoint
ep = aws.ec2.VpcEndpoint("ep",
    service_name=s3.service_name,
    vpc_id=foo.id)

