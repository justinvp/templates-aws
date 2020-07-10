import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myDevelopers = new aws.iam.Group("my_developers", {
    path: "/users/",
});
const myDeveloperPolicy = new aws.iam.GroupPolicy("my_developer_policy", {
    group: myDevelopers.id,
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

