import pulumi
import pulumi_aws as aws

a_ws_cloud_formation_stack_set_administration_role_assume_role_policy = aws.iam.get_policy_document(statements=[{
    "actions": ["sts:AssumeRole"],
    "effect": "Allow",
    "principals": [{
        "identifiers": ["cloudformation.amazonaws.com"],
        "type": "Service",
    }],
}])
a_ws_cloud_formation_stack_set_administration_role = aws.iam.Role("aWSCloudFormationStackSetAdministrationRole", assume_role_policy=a_ws_cloud_formation_stack_set_administration_role_assume_role_policy.json)
example = aws.cloudformation.StackSet("example",
    administration_role_arn=a_ws_cloud_formation_stack_set_administration_role.arn,
    parameters={
        "VPCCidr": "10.0.0.0/16",
    },
    template_body="""{
  "Parameters" : {
    "VPCCidr" : {
      "Type" : "String",
      "Default" : "10.0.0.0/16",
      "Description" : "Enter the CIDR block for the VPC. Default is 10.0.0.0/16."
    }
  },
  "Resources" : {
    "myVpc": {
      "Type" : "AWS::EC2::VPC",
      "Properties" : {
        "CidrBlock" : { "Ref" : "VPCCidr" },
        "Tags" : [
          {"Key": "Name", "Value": "Primary_CF_VPC"}
        ]
      }
    }
  }
}

""")
a_ws_cloud_formation_stack_set_administration_role_execution_policy_policy_document = example.execution_role_name.apply(lambda execution_role_name: aws.iam.get_policy_document(statements=[{
    "actions": ["sts:AssumeRole"],
    "effect": "Allow",
    "resources": [f"arn:aws:iam::*:role/{execution_role_name}"],
}]))
a_ws_cloud_formation_stack_set_administration_role_execution_policy_role_policy = aws.iam.RolePolicy("aWSCloudFormationStackSetAdministrationRoleExecutionPolicyRolePolicy",
    policy=a_ws_cloud_formation_stack_set_administration_role_execution_policy_policy_document.json,
    role=a_ws_cloud_formation_stack_set_administration_role.name)

