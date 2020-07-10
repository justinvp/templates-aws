using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var sample_ledger = new Aws.Qldb.Ledger("sample-ledger", new Aws.Qldb.LedgerArgs
        {
        });
    }

}

