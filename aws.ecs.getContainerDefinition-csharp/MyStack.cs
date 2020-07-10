using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ecs_mongo = Output.Create(Aws.Ecs.GetContainerDefinition.InvokeAsync(new Aws.Ecs.GetContainerDefinitionArgs
        {
            ContainerName = "mongodb",
            TaskDefinition = aws_ecs_task_definition.Mongo.Id,
        }));
    }

}

