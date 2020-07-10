import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const account = new aws.organizations.PolicyAttachment("account", {
    policyId: aws_organizations_policy_example.id,
    targetId: "123456789012",
});

