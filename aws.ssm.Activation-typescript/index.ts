import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testRole = new aws.iam.Role("test_role", {
    assumeRolePolicy: `  {
    "Version": "2012-10-17",
    "Statement": {
      "Effect": "Allow",
      "Principal": {"Service": "ssm.amazonaws.com"},
      "Action": "sts:AssumeRole"
    }
  }
`,
});
const testAttach = new aws.iam.RolePolicyAttachment("test_attach", {
    policyArn: "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore",
    role: testRole.name,
});
const foo = new aws.ssm.Activation("foo", {
    description: "Test",
    iamRole: testRole.id,
    registrationLimit: 5,
}, { dependsOn: [testAttach] });

