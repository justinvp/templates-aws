using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var current = Output.Create(Aws.Iam.GetAccountAlias.InvokeAsync());
        this.AccountId = current.Apply(current => current.AccountAlias);
    }

    [Output("accountId")]
    public Output<string> AccountId { get; set; }
}

