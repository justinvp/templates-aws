using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ecs_mongo = Output.Create(Aws.Ecs.GetCluster.InvokeAsync(new Aws.Ecs.GetClusterArgs
        {
            ClusterName = "ecs-mongo-production",
        }));
    }

}

