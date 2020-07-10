import pulumi
import pulumi_aws as aws

example_application = aws.codedeploy.Application("exampleApplication")
example_deployment_group = aws.codedeploy.DeploymentGroup("exampleDeploymentGroup",
    app_name=example_application.name,
    blue_green_deployment_config={
        "deploymentReadyOption": {
            "actionOnTimeout": "STOP_DEPLOYMENT",
            "waitTimeInMinutes": 60,
        },
        "greenFleetProvisioningOption": {
            "action": "DISCOVER_EXISTING",
        },
        "terminateBlueInstancesOnDeploymentSuccess": {
            "action": "KEEP_ALIVE",
        },
    },
    deployment_group_name="example-group",
    deployment_style={
        "deploymentOption": "WITH_TRAFFIC_CONTROL",
        "deploymentType": "BLUE_GREEN",
    },
    load_balancer_info={
        "elbInfo": [{
            "name": aws_elb["example"]["name"],
        }],
    },
    service_role_arn=aws_iam_role["example"]["arn"])

