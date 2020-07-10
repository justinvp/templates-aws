import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Declare the data source
const s3 = aws_vpc_foo.id.apply(id => aws.ec2.getVpcEndpoint({
    serviceName: "com.amazonaws.us-west-2.s3",
    vpcId: id,
}, { async: true }));
const privateS3 = new aws.ec2.VpcEndpointRouteTableAssociation("private_s3", {
    routeTableId: aws_route_table_private.id,
    vpcEndpointId: s3.id!,
});

