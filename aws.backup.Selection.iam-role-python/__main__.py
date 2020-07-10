import pulumi
import pulumi_aws as aws

example_role = aws.iam.Role("exampleRole", assume_role_policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": ["sts:AssumeRole"],
      "Effect": "allow",
      "Principal": {
        "Service": ["backup.amazonaws.com"]
      }
    }
  ]
}

""")
example_role_policy_attachment = aws.iam.RolePolicyAttachment("exampleRolePolicyAttachment",
    policy_arn="arn:aws:iam::aws:policy/service-role/AWSBackupServiceRolePolicyForBackup",
    role=example_role.name)
example_selection = aws.backup.Selection("exampleSelection", iam_role_arn=example_role.arn)

