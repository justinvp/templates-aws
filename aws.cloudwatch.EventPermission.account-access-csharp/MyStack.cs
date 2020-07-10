using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var devAccountAccess = new Aws.CloudWatch.EventPermission("devAccountAccess", new Aws.CloudWatch.EventPermissionArgs
        {
            Principal = "123456789012",
            StatementId = "DevAccountAccess",
        });
    }

}

