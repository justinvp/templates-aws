using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var multiDocker = Output.Create(Aws.ElasticBeanstalk.GetSolutionStack.InvokeAsync(new Aws.ElasticBeanstalk.GetSolutionStackArgs
        {
            MostRecent = true,
            NameRegex = "^64bit Amazon Linux (.*) Multi-container Docker (.*)$",
        }));
    }

}

