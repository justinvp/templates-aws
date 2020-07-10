import pulumi
import pulumi_aws as aws

my_api_gateway_vpc_link = aws.apigateway.get_vpc_link(name="my-vpc-link")

