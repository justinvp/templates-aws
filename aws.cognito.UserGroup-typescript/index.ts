import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const mainUserPool = new aws.cognito.UserPool("main", {});
const groupRole = new aws.iam.Role("group_role", {
    assumeRolePolicy: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Federated": "cognito-identity.amazonaws.com"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "cognito-identity.amazonaws.com:aud": "us-east-1:12345678-dead-beef-cafe-123456790ab"
        },
        "ForAnyValue:StringLike": {
          "cognito-identity.amazonaws.com:amr": "authenticated"
        }
      }
    }
  ]
}
`,
});
const mainUserGroup = new aws.cognito.UserGroup("main", {
    description: "Managed by Pulumi",
    precedence: 42,
    roleArn: groupRole.arn,
    userPoolId: mainUserPool.id,
});

