using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var serviceImage = Output.Create(Aws.Ecr.GetImage.InvokeAsync(new Aws.Ecr.GetImageArgs
        {
            ImageTag = "latest",
            RepositoryName = "my/service",
        }));
    }

}

