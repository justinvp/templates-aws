using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var service = Output.Create(Aws.Ecr.GetRepository.InvokeAsync(new Aws.Ecr.GetRepositoryArgs
        {
            Name = "ecr-repository",
        }));
    }

}

