import pulumi
import pulumi_aws as aws

test = aws.ec2.get_vpc_endpoint_service(filters=[{
    "name": "service-name",
    "values": ["some-service"],
}])

