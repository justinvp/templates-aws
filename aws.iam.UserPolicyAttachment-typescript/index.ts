import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const user = new aws.iam.User("user", {});
const policy = new aws.iam.Policy("policy", {
    description: "A test policy",
    policy: "", // insert policy here
});
const test_attach = new aws.iam.UserPolicyAttachment("test-attach", {
    policyArn: policy.arn,
    user: user.name,
});

