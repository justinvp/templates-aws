using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = Output.Create(Aws.CodeCommit.GetRepository.InvokeAsync(new Aws.CodeCommit.GetRepositoryArgs
        {
            RepositoryName = "MyTestRepository",
        }));
    }

}

