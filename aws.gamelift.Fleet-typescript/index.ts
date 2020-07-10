import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.gamelift.Fleet("example", {
    buildId: aws_gamelift_build_example.id,
    ec2InstanceType: "t2.micro",
    fleetType: "ON_DEMAND",
    runtimeConfiguration: {
        serverProcesses: [{
            concurrentExecutions: 1,
            launchPath: "C:\\game\\GomokuServer.exe",
        }],
    },
});

