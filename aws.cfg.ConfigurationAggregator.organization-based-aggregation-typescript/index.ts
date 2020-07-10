import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const organizationRole = new aws.iam.Role("organization", {
    assumeRolePolicy: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "config.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
`,
});
const organizationRolePolicyAttachment = new aws.iam.RolePolicyAttachment("organization", {
    policyArn: "arn:aws:iam::aws:policy/service-role/AWSConfigRoleForOrganizations",
    role: organizationRole.name,
});
const organizationConfigurationAggregator = new aws.cfg.ConfigurationAggregator("organization", {
    organizationAggregationSource: {
        allRegions: true,
        roleArn: organizationRole.arn,
    },
}, { dependsOn: [organizationRolePolicyAttachment] });

