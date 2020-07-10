import pulumi
import pulumi_aws as aws

a_ws_cloud_formation_stack_set_execution_role_assume_role_policy = aws.iam.get_policy_document(statements=[{
    "actions": ["sts:AssumeRole"],
    "effect": "Allow",
    "principals": [{
        "identifiers": [aws_iam_role["AWSCloudFormationStackSetAdministrationRole"]["arn"]],
        "type": "AWS",
    }],
}])
a_ws_cloud_formation_stack_set_execution_role = aws.iam.Role("aWSCloudFormationStackSetExecutionRole", assume_role_policy=a_ws_cloud_formation_stack_set_execution_role_assume_role_policy.json)
a_ws_cloud_formation_stack_set_execution_role_minimum_execution_policy_policy_document = aws.iam.get_policy_document(statements=[{
    "actions": [
        "cloudformation:*",
        "s3:*",
        "sns:*",
    ],
    "effect": "Allow",
    "resources": ["*"],
}])
a_ws_cloud_formation_stack_set_execution_role_minimum_execution_policy_role_policy = aws.iam.RolePolicy("aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyRolePolicy",
    policy=a_ws_cloud_formation_stack_set_execution_role_minimum_execution_policy_policy_document.json,
    role=a_ws_cloud_formation_stack_set_execution_role.name)

