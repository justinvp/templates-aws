import pulumi
import pulumi_aws as aws

example_application = aws.codedeploy.Application("exampleApplication", compute_platform="ECS")
example_deployment_group = aws.codedeploy.DeploymentGroup("exampleDeploymentGroup",
    app_name=example_application.name,
    auto_rollback_configuration={
        "enabled": True,
        "events": ["DEPLOYMENT_FAILURE"],
    },
    blue_green_deployment_config={
        "deploymentReadyOption": {
            "actionOnTimeout": "CONTINUE_DEPLOYMENT",
        },
        "terminateBlueInstancesOnDeploymentSuccess": {
            "action": "TERMINATE",
            "terminationWaitTimeInMinutes": 5,
        },
    },
    deployment_config_name="CodeDeployDefault.ECSAllAtOnce",
    deployment_group_name="example",
    deployment_style={
        "deploymentOption": "WITH_TRAFFIC_CONTROL",
        "deploymentType": "BLUE_GREEN",
    },
    ecs_service={
        "cluster_name": aws_ecs_cluster["example"]["name"],
        "service_name": aws_ecs_service["example"]["name"],
    },
    load_balancer_info={
        "targetGroupPairInfo": {
            "prodTrafficRoute": {
                "listenerArns": [aws_lb_listener["example"]["arn"]],
            },
            "targetGroup": [
                {
                    "name": aws_lb_target_group["blue"]["name"],
                },
                {
                    "name": aws_lb_target_group["green"]["name"],
                },
            ],
        },
    },
    service_role_arn=aws_iam_role["example"]["arn"])

