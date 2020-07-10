import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const organizationAccess = new aws.cloudwatch.EventPermission("OrganizationAccess", {
    condition: {
        key: "aws:PrincipalOrgID",
        type: "StringEquals",
        value: aws_organizations_organization_example.id,
    },
    principal: "*",
    statementId: "OrganizationAccess",
});

