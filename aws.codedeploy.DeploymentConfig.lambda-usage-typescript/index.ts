import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fooDeploymentConfig = new aws.codedeploy.DeploymentConfig("foo", {
    computePlatform: "Lambda",
    deploymentConfigName: "test-deployment-config",
    trafficRoutingConfig: {
        timeBasedLinear: {
            interval: 10,
            percentage: 10,
        },
        type: "TimeBasedLinear",
    },
});
const fooDeploymentGroup = new aws.codedeploy.DeploymentGroup("foo", {
    alarmConfiguration: {
        alarms: ["my-alarm-name"],
        enabled: true,
    },
    appName: aws_codedeploy_app_foo_app.name,
    autoRollbackConfiguration: {
        enabled: true,
        events: ["DEPLOYMENT_STOP_ON_ALARM"],
    },
    deploymentConfigName: fooDeploymentConfig.id,
    deploymentGroupName: "bar",
    serviceRoleArn: aws_iam_role_foo_role.arn,
});

