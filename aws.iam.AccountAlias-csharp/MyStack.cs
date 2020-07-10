using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @alias = new Aws.Iam.AccountAlias("alias", new Aws.Iam.AccountAliasArgs
        {
            AccountAlias = "my-account-alias",
        });
    }

}

