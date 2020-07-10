import pulumi
import pulumi_aws as aws

example_permission = aws.lambda_.Permission("examplePermission",
    action="lambda:InvokeFunction",
    function=aws_lambda_function["example"]["arn"],
    principal="config.amazonaws.com")
example_organization = aws.organizations.Organization("exampleOrganization",
    aws_service_access_principals=["config-multiaccountsetup.amazonaws.com"],
    feature_set="ALL")
example_organization_custom_rule = aws.cfg.OrganizationCustomRule("exampleOrganizationCustomRule",
    lambda_function_arn=aws_lambda_function["example"]["arn"],
    trigger_types=["ConfigurationItemChangeNotification"],
    opts=ResourceOptions(depends_on=[
            "aws_lambda_permission.example",
            "aws_organizations_organization.example",
        ]))

