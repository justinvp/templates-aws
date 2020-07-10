import pulumi
import pulumi_aws as aws

example_organization = aws.organizations.Organization("exampleOrganization",
    aws_service_access_principals=["config-multiaccountsetup.amazonaws.com"],
    feature_set="ALL")
example_organization_managed_rule = aws.cfg.OrganizationManagedRule("exampleOrganizationManagedRule", rule_identifier="IAM_PASSWORD_POLICY",
opts=ResourceOptions(depends_on=["aws_organizations_organization.example"]))

