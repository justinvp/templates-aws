using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = Output.Create(Aws.Ssm.GetDocument.InvokeAsync(new Aws.Ssm.GetDocumentArgs
        {
            DocumentFormat = "YAML",
            Name = "AWS-GatherSoftwareInventory",
        }));
        this.Content = foo.Apply(foo => foo.Content);
    }

    [Output("content")]
    public Output<string> Content { get; set; }
}

