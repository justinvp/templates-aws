import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const lbUser = new aws.iam.User("lb", {
    path: "/system/",
    tags: {
        "tag-key": "tag-value",
    },
});
const lbAccessKey = new aws.iam.AccessKey("lb", {
    user: lbUser.name,
});
const lbRo = new aws.iam.UserPolicy("lb_ro", {
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
    user: lbUser.name,
});

