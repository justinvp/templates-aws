using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var yada = new Aws.CloudWatch.LogGroup("yada", new Aws.CloudWatch.LogGroupArgs
        {
            Tags = 
            {
                { "Application", "serviceA" },
                { "Environment", "production" },
            },
        });
    }

}

