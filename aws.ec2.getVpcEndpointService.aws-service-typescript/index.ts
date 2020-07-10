import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Declare the data source
const s3 = pulumi.output(aws.ec2.getVpcEndpointService({
    service: "s3",
}, { async: true }));
// Create a VPC
const foo = new aws.ec2.Vpc("foo", {
    cidrBlock: "10.0.0.0/16",
});
// Create a VPC endpoint
const ep = new aws.ec2.VpcEndpoint("ep", {
    serviceName: s3.serviceName!,
    vpcId: foo.id,
});

