using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleEventCategories = Output.Create(Aws.Rds.GetEventCategories.InvokeAsync());
        this.Example = exampleEventCategories.Apply(exampleEventCategories => exampleEventCategories.EventCategories);
    }

    [Output("example")]
    public Output<string> Example { get; set; }
}

