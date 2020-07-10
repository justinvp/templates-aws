import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleRole = new aws.iam.Role("example", {
    assumeRolePolicy: `{
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
`,
});
const aWSCodeDeployRole = new aws.iam.RolePolicyAttachment("AWSCodeDeployRole", {
    policyArn: "arn:aws:iam::aws:policy/service-role/AWSCodeDeployRole",
    role: exampleRole.name,
});
const exampleApplication = new aws.codedeploy.Application("example", {});
const exampleTopic = new aws.sns.Topic("example", {});
const exampleDeploymentGroup = new aws.codedeploy.DeploymentGroup("example", {
    alarmConfiguration: {
        alarms: ["my-alarm-name"],
        enabled: true,
    },
    appName: exampleApplication.name,
    autoRollbackConfiguration: {
        enabled: true,
        events: ["DEPLOYMENT_FAILURE"],
    },
    deploymentGroupName: "example-group",
    ec2TagSets: [{
        ec2TagFilters: [
            {
                key: "filterkey1",
                type: "KEY_AND_VALUE",
                value: "filtervalue",
            },
            {
                key: "filterkey2",
                type: "KEY_AND_VALUE",
                value: "filtervalue",
            },
        ],
    }],
    serviceRoleArn: exampleRole.arn,
    triggerConfigurations: [{
        triggerEvents: ["DeploymentFailure"],
        triggerName: "example-trigger",
        triggerTargetArn: exampleTopic.arn,
    }],
});

