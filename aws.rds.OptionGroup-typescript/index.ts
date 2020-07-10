import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.rds.OptionGroup("example", {
    engineName: "sqlserver-ee",
    majorEngineVersion: "11.00",
    options: [
        {
            optionName: "Timezone",
            optionSettings: [{
                name: "TIME_ZONE",
                value: "UTC",
            }],
        },
        {
            optionName: "SQLSERVER_BACKUP_RESTORE",
            optionSettings: [{
                name: "IAM_ROLE_ARN",
                value: aws_iam_role_example.arn,
            }],
        },
        {
            optionName: "TDE",
        },
    ],
    optionGroupDescription: "Option Group",
});

