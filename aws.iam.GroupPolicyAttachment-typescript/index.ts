import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const group = new aws.iam.Group("group", {});
const policy = new aws.iam.Policy("policy", {
    description: "A test policy",
    policy: "", // insert policy here
});
const test_attach = new aws.iam.GroupPolicyAttachment("test-attach", {
    group: group.name,
    policyArn: policy.arn,
});

