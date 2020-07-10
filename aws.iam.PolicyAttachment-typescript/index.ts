import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const user = new aws.iam.User("user", {});
const role = new aws.iam.Role("role", {
    assumeRolePolicy: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
`,
});
const group = new aws.iam.Group("group", {});
const policy = new aws.iam.Policy("policy", {
    description: "A test policy",
    policy: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "ec2:Describe*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
`,
});
const test_attach = new aws.iam.PolicyAttachment("test-attach", {
    groups: [group.name],
    policyArn: policy.arn,
    roles: [role.name],
    users: [user.name],
});

