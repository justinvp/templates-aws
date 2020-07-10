import pulumi
import pulumi_aws as aws

example_organization = aws.organizations.Organization("exampleOrganization",
    aws_service_access_principals=["guardduty.amazonaws.com"],
    feature_set="ALL")
example_detector = aws.guardduty.Detector("exampleDetector")
example_organization_admin_account = aws.guardduty.OrganizationAdminAccount("exampleOrganizationAdminAccount", admin_account_id="123456789012",
opts=ResourceOptions(depends_on=[example_organization]))

