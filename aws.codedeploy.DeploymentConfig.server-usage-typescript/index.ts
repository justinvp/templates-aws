import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fooDeploymentConfig = new aws.codedeploy.DeploymentConfig("foo", {
    deploymentConfigName: "test-deployment-config",
    minimumHealthyHosts: {
        type: "HOST_COUNT",
        value: 2,
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
        events: ["DEPLOYMENT_FAILURE"],
    },
    deploymentConfigName: fooDeploymentConfig.id,
    deploymentGroupName: "bar",
    ec2TagFilters: [{
        key: "filterkey",
        type: "KEY_AND_VALUE",
        value: "filtervalue",
    }],
    serviceRoleArn: aws_iam_role_foo_role.arn,
    triggerConfigurations: [{
        triggerEvents: ["DeploymentFailure"],
        triggerName: "foo-trigger",
        triggerTargetArn: "foo-topic-arn",
    }],
});

