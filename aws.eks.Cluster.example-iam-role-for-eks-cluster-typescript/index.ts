import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.iam.Role("example", {
    assumeRolePolicy: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "eks.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
`,
});
const example_AmazonEKSClusterPolicy = new aws.iam.RolePolicyAttachment("example-AmazonEKSClusterPolicy", {
    policyArn: "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy",
    role: example.name,
});
const example_AmazonEKSServicePolicy = new aws.iam.RolePolicyAttachment("example-AmazonEKSServicePolicy", {
    policyArn: "arn:aws:iam::aws:policy/AmazonEKSServicePolicy",
    role: example.name,
});

