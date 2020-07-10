import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.datasync.Agent("example", {
    ipAddress: "1.2.3.4",
});

