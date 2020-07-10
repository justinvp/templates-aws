import pulumi
import pulumi_aws as aws

foo_deployment_config = aws.codedeploy.DeploymentConfig("fooDeploymentConfig",
    deployment_config_name="test-deployment-config",
    minimum_healthy_hosts={
        "type": "HOST_COUNT",
        "value": 2,
    })
foo_deployment_group = aws.codedeploy.DeploymentGroup("fooDeploymentGroup",
    alarm_configuration={
        "alarms": ["my-alarm-name"],
        "enabled": True,
    },
    app_name=aws_codedeploy_app["foo_app"]["name"],
    auto_rollback_configuration={
        "enabled": True,
        "events": ["DEPLOYMENT_FAILURE"],
    },
    deployment_config_name=foo_deployment_config.id,
    deployment_group_name="bar",
    ec2_tag_filters=[{
        "key": "filterkey",
        "type": "KEY_AND_VALUE",
        "value": "filtervalue",
    }],
    service_role_arn=aws_iam_role["foo_role"]["arn"],
    trigger_configurations=[{
        "triggerEvents": ["DeploymentFailure"],
        "triggerName": "foo-trigger",
        "triggerTargetArn": "foo-topic-arn",
    }])

