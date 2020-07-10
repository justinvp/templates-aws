import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const network = new aws.cloudformation.Stack("network", {
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

