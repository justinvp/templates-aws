import pulumi
import pulumi_aws as aws

main_route_table_association = aws.ec2.MainRouteTableAssociation("mainRouteTableAssociation",
    route_table_id=aws_route_table["bar"]["id"],
    vpc_id=aws_vpc["foo"]["id"])

