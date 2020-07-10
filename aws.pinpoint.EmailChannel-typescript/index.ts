import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const app = new aws.pinpoint.App("app", {});
const identity = new aws.ses.DomainIdentity("identity", {
    domain: "example.com",
});
const role = new aws.iam.Role("role", {
    assumeRolePolicy: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "pinpoint.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
`,
});
const email = new aws.pinpoint.EmailChannel("email", {
    applicationId: app.applicationId,
    fromAddress: "user@example.com",
    identity: identity.arn,
    roleArn: role.arn,
});
const rolePolicy = new aws.iam.RolePolicy("role_policy", {
    policy: `{
  "Version": "2012-10-17",
  "Statement": {
    "Action": [
      "mobileanalytics:PutEvents",
      "mobileanalytics:PutItems"
    ],
    "Effect": "Allow",
    "Resource": [
      "*"
    ]
  }
}
`,
    role: role.id,
});

