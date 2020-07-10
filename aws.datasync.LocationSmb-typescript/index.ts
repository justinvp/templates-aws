import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.datasync.LocationSmb("example", {
    agentArns: [aws_datasync_agent_example.arn],
    password: "ANotGreatPassword",
    serverHostname: "smb.example.com",
    subdirectory: "/exported/path",
    user: "Guest",
});

