import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const contractors = new aws.workspaces.IpGroup("contractors", {
    description: "Contractors IP access control group",
});

