import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleRole = new aws.iam.Role("example", {
    assumeRolePolicy: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": ["sts:AssumeRole"],
      "Effect": "allow",
      "Principal": {
        "Service": ["backup.amazonaws.com"]
      }
    }
  ]
}
`,
});
const exampleRolePolicyAttachment = new aws.iam.RolePolicyAttachment("example", {
    policyArn: "arn:aws:iam::aws:policy/service-role/AWSBackupServiceRolePolicyForBackup",
    role: exampleRole.name,
});
const exampleSelection = new aws.backup.Selection("example", {
    iamRoleArn: exampleRole.arn,
});

