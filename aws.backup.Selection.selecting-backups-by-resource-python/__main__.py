import pulumi
import pulumi_aws as aws

example = aws.backup.Selection("example",
    iam_role_arn=aws_iam_role["example"]["arn"],
    plan_id=aws_backup_plan["example"]["id"],
    resources=[
        aws_db_instance["example"]["arn"],
        aws_ebs_volume["example"]["arn"],
        aws_efs_file_system["example"]["arn"],
    ])

