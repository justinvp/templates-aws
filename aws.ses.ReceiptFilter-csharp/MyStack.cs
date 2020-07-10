using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var filter = new Aws.Ses.ReceiptFilter("filter", new Aws.Ses.ReceiptFilterArgs
        {
            Cidr = "10.10.10.10",
            Policy = "Block",
        });
    }

}

