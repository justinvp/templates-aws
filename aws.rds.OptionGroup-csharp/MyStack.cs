using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Rds.OptionGroup("example", new Aws.Rds.OptionGroupArgs
        {
            EngineName = "sqlserver-ee",
            MajorEngineVersion = "11.00",
            Options = 
            {
                new Aws.Rds.Inputs.OptionGroupOptionArgs
                {
                    OptionName = "Timezone",
                    OptionSettings = 
                    {
                        new Aws.Rds.Inputs.OptionGroupOptionOptionSettingArgs
                        {
                            Name = "TIME_ZONE",
                            Value = "UTC",
                        },
                    },
                },
                new Aws.Rds.Inputs.OptionGroupOptionArgs
                {
                    OptionName = "SQLSERVER_BACKUP_RESTORE",
                    OptionSettings = 
                    {
                        new Aws.Rds.Inputs.OptionGroupOptionOptionSettingArgs
                        {
                            Name = "IAM_ROLE_ARN",
                            Value = aws_iam_role.Example.Arn,
                        },
                    },
                },
                new Aws.Rds.Inputs.OptionGroupOptionArgs
                {
                    OptionName = "TDE",
                },
            },
            OptionGroupDescription = "Option Group",
        });
    }

}

