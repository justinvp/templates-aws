import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const root = new aws.organizations.PolicyAttachment("root", {
    policyId: aws_organizations_policy_example.id,
    targetId: aws_organizations_organization_example.roots.0.id,
});

