import pulumi
import pulumi_aws as aws

organization_role = aws.iam.Role("organizationRole", assume_role_policy="""{
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

""")
organization_configuration_aggregator = aws.cfg.ConfigurationAggregator("organizationConfigurationAggregator", organization_aggregation_source={
    "allRegions": True,
    "role_arn": organization_role.arn,
},
opts=ResourceOptions(depends_on=["aws_iam_role_policy_attachment.organization"]))
organization_role_policy_attachment = aws.iam.RolePolicyAttachment("organizationRolePolicyAttachment",
    policy_arn="arn:aws:iam::aws:policy/service-role/AWSConfigRoleForOrganizations",
    role=organization_role.name)

