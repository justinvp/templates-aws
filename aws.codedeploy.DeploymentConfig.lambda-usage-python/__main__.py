import pulumi
import pulumi_aws as aws

foo_deployment_config = aws.codedeploy.DeploymentConfig("fooDeploymentConfig",
    compute_platform="Lambda",
    deployment_config_name="test-deployment-config",
    traffic_routing_config={
        "timeBasedLinear": {
            "interval": 10,
            "percentage": 10,
        },
        "type": "TimeBasedLinear",
    })
foo_deployment_group = aws.codedeploy.DeploymentGroup("fooDeploymentGroup",
    alarm_configuration={
        "alarms": ["my-alarm-name"],
        "enabled": True,
    },
    app_name=aws_codedeploy_app["foo_app"]["name"],
    auto_rollback_configuration={
        "enabled": True,
        "events": ["DEPLOYMENT_STOP_ON_ALARM"],
    },
    deployment_config_name=foo_deployment_config.id,
    deployment_group_name="bar",
    service_role_arn=aws_iam_role["foo_role"]["arn"])

