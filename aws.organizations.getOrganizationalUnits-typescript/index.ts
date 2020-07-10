import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const org = aws.organizations.getOrganization({});
const ou = org.then(org => aws.organizations.getOrganizationalUnits({
    parentId: org.roots[0].id,
}));

