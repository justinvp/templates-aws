import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultSubnetGroup = new aws.neptune.SubnetGroup("default", {
    subnetIds: [
        aws_subnet_frontend.id,
        aws_subnet_backend.id,
    ],
    tags: {
        Name: "My neptune subnet group",
    },
});

