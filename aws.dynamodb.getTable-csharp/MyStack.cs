using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var tableName = Output.Create(Aws.DynamoDB.GetTable.InvokeAsync(new Aws.DynamoDB.GetTableArgs
        {
            Name = "tableName",
        }));
    }

}

