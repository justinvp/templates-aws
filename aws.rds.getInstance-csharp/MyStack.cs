using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var database = Output.Create(Aws.Rds.GetInstance.InvokeAsync(new Aws.Rds.GetInstanceArgs
        {
            DbInstanceIdentifier = "my-test-database",
        }));
    }

}

