import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleOrganization = new aws.organizations.Organization("exampleOrganization", {
    awsServiceAccessPrincipals: ["guardduty.amazonaws.com"],
    featureSet: "ALL",
});
const exampleDetector = new aws.guardduty.Detector("exampleDetector", {});
const exampleOrganizationAdminAccount = new aws.guardduty.OrganizationAdminAccount("exampleOrganizationAdminAccount", {adminAccountId: "123456789012"}, {
    dependsOn: [exampleOrganization],
});

