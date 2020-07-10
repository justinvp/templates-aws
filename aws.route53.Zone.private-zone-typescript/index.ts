import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const privateZone = new aws.route53.Zone("private", {
    vpcs: [{
        vpcId: aws_vpc_example.id,
    }],
});

