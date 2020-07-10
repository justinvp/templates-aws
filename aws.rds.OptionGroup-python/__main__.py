import pulumi
import pulumi_aws as aws

example = aws.rds.OptionGroup("example",
    engine_name="sqlserver-ee",
    major_engine_version="11.00",
    options=[
        {
            "optionName": "Timezone",
            "optionSettings": [{
                "name": "TIME_ZONE",
                "value": "UTC",
            }],
        },
        {
            "optionName": "SQLSERVER_BACKUP_RESTORE",
            "optionSettings": [{
                "name": "IAM_ROLE_ARN",
                "value": aws_iam_role["example"]["arn"],
            }],
        },
        {
            "optionName": "TDE",
        },
    ],
    option_group_description="Option Group")

