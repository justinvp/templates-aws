import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleApplication = new aws.codedeploy.Application("example", {});
const exampleDeploymentGroup = new aws.codedeploy.DeploymentGroup("example", {
    appName: exampleApplication.name,
    blueGreenDeploymentConfig: {
        deploymentReadyOption: {
            actionOnTimeout: "STOP_DEPLOYMENT",
            waitTimeInMinutes: 60,
        },
        greenFleetProvisioningOption: {
            action: "DISCOVER_EXISTING",
        },
        terminateBlueInstancesOnDeploymentSuccess: {
            action: "KEEP_ALIVE",
        },
    },
    deploymentGroupName: "example-group",
    deploymentStyle: {
        deploymentOption: "WITH_TRAFFIC_CONTROL",
        deploymentType: "BLUE_GREEN",
    },
    loadBalancerInfo: {
        elbInfos: [{
            name: aws_elb_example.name,
        }],
    },
    serviceRoleArn: aws_iam_role_example.arn,
});

