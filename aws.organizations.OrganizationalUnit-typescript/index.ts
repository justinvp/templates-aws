import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.organizations.OrganizationalUnit("example", {
    parentId: aws_organizations_organization_example.roots.0.id,
});

