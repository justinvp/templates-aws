using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Qldb.GetLedger.InvokeAsync(new Aws.Qldb.GetLedgerArgs
        {
            Name = "an_example_ledger",
        }));
    }

}

