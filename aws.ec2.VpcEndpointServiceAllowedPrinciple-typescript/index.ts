import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = pulumi.output(aws.getCallerIdentity({ async: true }));
const allowMeToFoo = new aws.ec2.VpcEndpointServiceAllowedPrinciple("allow_me_to_foo", {
    principalArn: current.arn,
    vpcEndpointServiceId: aws_vpc_endpoint_service_foo.id,
});

