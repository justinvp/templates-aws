import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const key = new aws.kms.Key("a", {});
const role = new aws.iam.Role("a", {
    assumeRolePolicy: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
`,
});
const grant = new aws.kms.Grant("a", {
    constraints: [{
        encryptionContextEquals: {
            Department: "Finance",
        },
    }],
    granteePrincipal: role.arn,
    keyId: key.keyId,
    operations: [
        "Encrypt",
        "Decrypt",
        "GenerateDataKey",
    ],
});

