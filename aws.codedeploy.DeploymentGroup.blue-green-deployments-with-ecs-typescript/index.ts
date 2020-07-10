import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleApplication = new aws.codedeploy.Application("example", {
    computePlatform: "ECS",
});
const exampleDeploymentGroup = new aws.codedeploy.DeploymentGroup("example", {
    appName: exampleApplication.name,
    autoRollbackConfiguration: {
        enabled: true,
        events: ["DEPLOYMENT_FAILURE"],
    },
    blueGreenDeploymentConfig: {
        deploymentReadyOption: {
            actionOnTimeout: "CONTINUE_DEPLOYMENT",
        },
        terminateBlueInstancesOnDeploymentSuccess: {
            action: "TERMINATE",
            terminationWaitTimeInMinutes: 5,
        },
    },
    deploymentConfigName: "CodeDeployDefault.ECSAllAtOnce",
    deploymentGroupName: "example",
    deploymentStyle: {
        deploymentOption: "WITH_TRAFFIC_CONTROL",
        deploymentType: "BLUE_GREEN",
    },
    ecsService: {
        clusterName: aws_ecs_cluster_example.name,
        serviceName: aws_ecs_service_example.name,
    },
    loadBalancerInfo: {
        targetGroupPairInfo: {
            prodTrafficRoute: {
                listenerArns: [aws_lb_listener_example.arn],
            },
            targetGroups: [
                {
                    name: aws_lb_target_group_blue.name,
                },
                {
                    name: aws_lb_target_group_green.name,
                },
            ],
        },
    },
    serviceRoleArn: aws_iam_role_example.arn,
});

