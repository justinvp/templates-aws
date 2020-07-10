import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplePermission = new aws.lambda.Permission("example", {
    action: "lambda:InvokeFunction",
    function: aws_lambda_function_example.arn,
    principal: "config.amazonaws.com",
});
const exampleOrganization = new aws.organizations.Organization("example", {
    awsServiceAccessPrincipals: ["config-multiaccountsetup.amazonaws.com"],
    featureSet: "ALL",
});
const exampleOrganizationCustomRule = new aws.cfg.OrganizationCustomRule("example", {
    lambdaFunctionArn: aws_lambda_function_example.arn,
    triggerTypes: ["ConfigurationItemChangeNotification"],
}, { dependsOn: [examplePermission, exampleOrganization] });

