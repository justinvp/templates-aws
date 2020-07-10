using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bar = new Aws.Ecs.Service("bar", new Aws.Ecs.ServiceArgs
        {
            Cluster = aws_ecs_cluster.Foo.Id,
            SchedulingStrategy = "DAEMON",
            TaskDefinition = aws_ecs_task_definition.Bar.Arn,
        });
    }

}

