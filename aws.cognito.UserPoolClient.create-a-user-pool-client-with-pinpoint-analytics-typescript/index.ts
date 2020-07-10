import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = pulumi.output(aws.getCallerIdentity({ async: true }));
const testUserPool = new aws.cognito.UserPool("test", {});
const testApp = new aws.pinpoint.App("test", {});
const testRole = new aws.iam.Role("test", {
    assumeRolePolicy: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "cognito-idp.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
`,
});
const testRolePolicy = new aws.iam.RolePolicy("test", {
    policy: pulumi.interpolate`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "mobiletargeting:UpdateEndpoint",
        "mobiletargeting:PutItems"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:mobiletargeting:*:${current.accountId}:apps/${testApp.applicationId}*"
    }
  ]
}
`,
    role: testRole.id,
});
const testUserPoolClient = new aws.cognito.UserPoolClient("test", {
    analyticsConfiguration: {
        applicationId: testApp.applicationId,
        externalId: "some_id",
        roleArn: testRole.arn,
        userDataShared: true,
    },
    userPoolId: testUserPool.id,
});

