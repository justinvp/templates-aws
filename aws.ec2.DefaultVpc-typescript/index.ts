import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultDefaultVpc = new aws.ec2.DefaultVpc("default", {
    tags: {
        Name: "Default VPC",
    },
});

