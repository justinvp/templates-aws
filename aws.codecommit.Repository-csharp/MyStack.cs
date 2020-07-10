using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = new Aws.CodeCommit.Repository("test", new Aws.CodeCommit.RepositoryArgs
        {
            Description = "This is the Sample App Repository",
            RepositoryName = "MyTestRepository",
        });
    }

}

