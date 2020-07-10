using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Backup.Selection("example", new Aws.Backup.SelectionArgs
        {
            IamRoleArn = aws_iam_role.Example.Arn,
            PlanId = aws_backup_plan.Example.Id,
            Resources = 
            {
                aws_db_instance.Example.Arn,
                aws_ebs_volume.Example.Arn,
                aws_efs_file_system.Example.Arn,
            },
        });
    }

}

