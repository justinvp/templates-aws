using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.ElasticBeanstalk.GetApplication.InvokeAsync(new Aws.ElasticBeanstalk.GetApplicationArgs
        {
            Name = "example",
        }));
        this.Arn = example.Apply(example => example.Arn);
        this.Description = example.Apply(example => example.Description);
    }

    [Output("arn")]
    public Output<string> Arn { get; set; }
    [Output("description")]
    public Output<string> Description { get; set; }
}

