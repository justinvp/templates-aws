using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var current = Output.Create(Aws.GetCanonicalUserId.InvokeAsync());
        this.CanonicalUserId = current.Apply(current => current.Id);
    }

    [Output("canonicalUserId")]
    public Output<string> CanonicalUserId { get; set; }
}

