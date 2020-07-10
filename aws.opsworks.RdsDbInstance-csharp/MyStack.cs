using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var myInstance = new Aws.OpsWorks.RdsDbInstance("myInstance", new Aws.OpsWorks.RdsDbInstanceArgs
        {
            DbPassword = "somePass",
            DbUser = "someUser",
            RdsDbInstanceArn = aws_db_instance.My_instance.Arn,
            StackId = aws_opsworks_stack.My_stack.Id,
        });
    }

}

