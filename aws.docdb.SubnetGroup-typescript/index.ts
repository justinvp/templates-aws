import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultSubnetGroup = new aws.docdb.SubnetGroup("default", {
    subnetIds: [
        aws_subnet_frontend.id,
        aws_subnet_backend.id,
    ],
    tags: {
        Name: "My docdb subnet group",
    },
});

