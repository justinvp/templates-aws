import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultAz1 = new aws.ec2.DefaultSubnet("default_az1", {
    availabilityZone: "us-west-2a",
    tags: {
        Name: "Default subnet for us-west-2a",
    },
});

