import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultSecurityGroup = new aws.rds.SecurityGroup("default", {
    ingress: [{
        cidr: "10.0.0.0/24",
    }],
});

