import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.datasync.NfsLocation("example", {
    onPremConfig: {
        agentArns: [aws_datasync_agent_example.arn],
    },
    serverHostname: "nfs.example.com",
    subdirectory: "/exported/path",
});

