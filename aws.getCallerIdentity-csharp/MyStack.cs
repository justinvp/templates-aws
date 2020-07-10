using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var current = Output.Create(Aws.GetCallerIdentity.InvokeAsync());
        this.AccountId = current.Apply(current => current.AccountId);
        this.CallerArn = current.Apply(current => current.Arn);
        this.CallerUser = current.Apply(current => current.UserId);
    }

    [Output("accountId")]
    public Output<string> AccountId { get; set; }
    [Output("callerArn")]
    public Output<string> CallerArn { get; set; }
    [Output("callerUser")]
    public Output<string> CallerUser { get; set; }
}

