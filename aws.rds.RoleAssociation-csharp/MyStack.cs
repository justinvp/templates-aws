using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Rds.RoleAssociation("example", new Aws.Rds.RoleAssociationArgs
        {
            DbInstanceIdentifier = aws_db_instance.Example.Id,
            FeatureName = "S3_INTEGRATION",
            RoleArn = aws_iam_role.Example.Id,
        });
    }

}

