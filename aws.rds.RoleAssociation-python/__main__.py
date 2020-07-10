import pulumi
import pulumi_aws as aws

example = aws.rds.RoleAssociation("example",
    db_instance_identifier=aws_db_instance["example"]["id"],
    feature_name="S3_INTEGRATION",
    role_arn=aws_iam_role["example"]["id"])

