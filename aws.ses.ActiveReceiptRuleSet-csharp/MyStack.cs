using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var main = new Aws.Ses.ActiveReceiptRuleSet("main", new Aws.Ses.ActiveReceiptRuleSetArgs
        {
            RuleSetName = "primary-rules",
        });
    }

}

