using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var main = new Aws.Ses.ReceiptRuleSet("main", new Aws.Ses.ReceiptRuleSetArgs
        {
            RuleSetName = "primary-rules",
        });
    }

}

