import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy = pulumi.output(aws.iam.getPolicyDocument({
    statements: [{
        actions: ["sts:AssumeRole"],
        effect: "Allow",
        principals: [{
            identifiers: ["cloudformation.amazonaws.com"],
            type: "Service",
        }],
    }],
}, { async: true }));
const aWSCloudFormationStackSetAdministrationRole = new aws.iam.Role("AWSCloudFormationStackSetAdministrationRole", {
    assumeRolePolicy: aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy.json,
});
const example = new aws.cloudformation.StackSet("example", {
    administrationRoleArn: aWSCloudFormationStackSetAdministrationRole.arn,
    parameters: {
        VPCCidr: "10.0.0.0/16",
    },
    templateBody: `{
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
`,
});
const aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument = example.executionRoleName.apply(executionRoleName => aws.iam.getPolicyDocument({
    statements: [{
        actions: ["sts:AssumeRole"],
        effect: "Allow",
        resources: [`arn:aws:iam::*:role/${executionRoleName}`],
    }],
}, { async: true }));
const aWSCloudFormationStackSetAdministrationRoleExecutionPolicyRolePolicy = new aws.iam.RolePolicy("AWSCloudFormationStackSetAdministrationRole_ExecutionPolicy", {
    policy: aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument.json,
    role: aWSCloudFormationStackSetAdministrationRole.name,
});

