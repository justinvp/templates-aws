import pulumi
import pulumi_aws as aws

test_role = aws.iam.Role("testRole", assume_role_policy="""  {
    "Version": "2012-10-17",
    "Statement": {
      "Effect": "Allow",
      "Principal": {"Service": "ssm.amazonaws.com"},
      "Action": "sts:AssumeRole"
    }
  }

""")
test_attach = aws.iam.RolePolicyAttachment("testAttach",
    policy_arn="arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore",
    role=test_role.name)
foo = aws.ssm.Activation("foo",
    description="Test",
    iam_role=test_role.id,
    registration_limit="5",
    opts=ResourceOptions(depends_on=["aws_iam_role_policy_attachment.test_attach"]))

