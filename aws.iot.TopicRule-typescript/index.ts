import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const mytopic = new aws.sns.Topic("mytopic", {});
const role = new aws.iam.Role("role", {
    assumeRolePolicy: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "iot.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
`,
});
const rule = new aws.iot.TopicRule("rule", {
    description: "Example rule",
    enabled: true,
    sns: {
        messageFormat: "RAW",
        roleArn: role.arn,
        targetArn: mytopic.arn,
    },
    sql: "SELECT * FROM 'topic/test'",
    sqlVersion: "2016-03-23",
});
const iamPolicyForLambda = new aws.iam.RolePolicy("iam_policy_for_lambda", {
    policy: pulumi.interpolate`{
  "Version": "2012-10-17",
  "Statement": [
    {
        "Effect": "Allow",
        "Action": [
            "sns:Publish"
        ],
        "Resource": "${mytopic.arn}"
    }
  ]
}
`,
    role: role.id,
});

