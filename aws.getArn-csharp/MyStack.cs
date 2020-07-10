using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var dbInstance = Output.Create(Aws.GetArn.InvokeAsync(new Aws.GetArnArgs
        {
            Arn = "arn:aws:rds:eu-west-1:123456789012:db:mysql-db",
        }));
    }

}

