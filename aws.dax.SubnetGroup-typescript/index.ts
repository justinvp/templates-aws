import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.dax.SubnetGroup("example", {
    subnetIds: [
        aws_subnet_example1.id,
        aws_subnet_example2.id,
    ],
});

