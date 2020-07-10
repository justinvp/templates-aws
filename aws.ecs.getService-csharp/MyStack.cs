using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Ecs.GetService.InvokeAsync(new Aws.Ecs.GetServiceArgs
        {
            ClusterArn = data.Aws_ecs_cluster.Example.Arn,
            ServiceName = "example",
        }));
    }

}

