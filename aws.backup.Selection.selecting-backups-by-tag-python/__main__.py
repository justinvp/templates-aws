import pulumi
import pulumi_aws as aws

example = aws.backup.Selection("example",
    iam_role_arn=aws_iam_role["example"]["arn"],
    plan_id=aws_backup_plan["example"]["id"],
    selection_tags=[{
        "key": "foo",
        "type": "STRINGEQUALS",
        "value": "bar",
    }])

