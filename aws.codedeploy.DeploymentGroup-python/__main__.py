import pulumi
import pulumi_aws as aws

example_role = aws.iam.Role("exampleRole", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "codedeploy.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}

""")
a_ws_code_deploy_role = aws.iam.RolePolicyAttachment("aWSCodeDeployRole",
    policy_arn="arn:aws:iam::aws:policy/service-role/AWSCodeDeployRole",
    role=example_role.name)
example_application = aws.codedeploy.Application("exampleApplication")
example_topic = aws.sns.Topic("exampleTopic")
example_deployment_group = aws.codedeploy.DeploymentGroup("exampleDeploymentGroup",
    alarm_configuration={
        "alarms": ["my-alarm-name"],
        "enabled": True,
    },
    app_name=example_application.name,
    auto_rollback_configuration={
        "enabled": True,
        "events": ["DEPLOYMENT_FAILURE"],
    },
    deployment_group_name="example-group",
    ec2_tag_sets=[{
        "ec2TagFilter": [
            {
                "key": "filterkey1",
                "type": "KEY_AND_VALUE",
                "value": "filtervalue",
            },
            {
                "key": "filterkey2",
                "type": "KEY_AND_VALUE",
                "value": "filtervalue",
            },
        ],
    }],
    service_role_arn=example_role.arn,
    trigger_configurations=[{
        "triggerEvents": ["DeploymentFailure"],
        "triggerName": "example-trigger",
        "triggerTargetArn": example_topic.arn,
    }])

