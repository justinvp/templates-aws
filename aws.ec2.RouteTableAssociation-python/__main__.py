import pulumi
import pulumi_aws as aws

route_table_association = aws.ec2.RouteTableAssociation("routeTableAssociation",
    subnet_id=aws_subnet["foo"]["id"],
    route_table_id=aws_route_table["bar"]["id"])

