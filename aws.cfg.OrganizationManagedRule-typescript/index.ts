import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleOrganization = new aws.organizations.Organization("example", {
    awsServiceAccessPrincipals: ["config-multiaccountsetup.amazonaws.com"],
    featureSet: "ALL",
});
const exampleOrganizationManagedRule = new aws.cfg.OrganizationManagedRule("example", {
    ruleIdentifier: "IAM_PASSWORD_POLICY",
}, { dependsOn: [exampleOrganization] });

