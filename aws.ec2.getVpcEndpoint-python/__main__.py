import pulumi
import pulumi_aws as aws

s3 = aws.ec2.get_vpc_endpoint(service_name="com.amazonaws.us-west-2.s3",
    vpc_id=aws_vpc["foo"]["id"])
private_s3 = aws.ec2.VpcEndpointRouteTableAssociation("privateS3",
    route_table_id=aws_route_table["private"]["id"],
    vpc_endpoint_id=s3.id)

