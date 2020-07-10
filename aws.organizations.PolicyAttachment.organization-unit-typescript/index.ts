import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const unit = new aws.organizations.PolicyAttachment("unit", {
    policyId: aws_organizations_policy_example.id,
    targetId: aws_organizations_organizational_unit_example.id,
});

