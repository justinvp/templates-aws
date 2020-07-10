import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy = aws_iam_role_AWSCloudFormationStackSetAdministrationRole.arn.apply(arn => aws.iam.getPolicyDocument({
    statements: [{
        actions: ["sts:AssumeRole"],
        effect: "Allow",
        principals: [{
            identifiers: [arn],
            type: "AWS",
        }],
    }],
}, { async: true }));
const aWSCloudFormationStackSetExecutionRole = new aws.iam.Role("AWSCloudFormationStackSetExecutionRole", {
    assumeRolePolicy: aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy.json,
});
// Documentation: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs.html
// Additional IAM permissions necessary depend on the resources defined in the StackSet template
const aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument = pulumi.output(aws.iam.getPolicyDocument({
    statements: [{
        actions: [
            "cloudformation:*",
            "s3:*",
            "sns:*",
        ],
        effect: "Allow",
        resources: ["*"],
    }],
}, { async: true }));
const aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyRolePolicy = new aws.iam.RolePolicy("AWSCloudFormationStackSetExecutionRole_MinimumExecutionPolicy", {
    policy: aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument.json,
    role: aWSCloudFormationStackSetExecutionRole.name,
});

