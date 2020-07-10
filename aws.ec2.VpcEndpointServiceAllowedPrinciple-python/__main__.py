import pulumi
import pulumi_aws as aws

current = aws.get_caller_identity()
allow_me_to_foo = aws.ec2.VpcEndpointServiceAllowedPrinciple("allowMeToFoo",
    principal_arn=current.arn,
    vpc_endpoint_service_id=aws_vpc_endpoint_service["foo"]["id"])

