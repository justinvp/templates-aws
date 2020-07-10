using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Outposts.GetOutpostInstanceTypes.InvokeAsync(new Aws.Outposts.GetOutpostInstanceTypesArgs
        {
            Arn = data.Aws_outposts_outpost.Example.Arn,
        }));
    }

}

