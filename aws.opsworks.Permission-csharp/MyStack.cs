using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var myStackPermission = new Aws.OpsWorks.Permission("myStackPermission", new Aws.OpsWorks.PermissionArgs
        {
            AllowSsh = true,
            AllowSudo = true,
            Level = "iam_only",
            StackId = aws_opsworks_stack.Stack.Id,
            UserArn = aws_iam_user.User.Arn,
        });
    }

}

