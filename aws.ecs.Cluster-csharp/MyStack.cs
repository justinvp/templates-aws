using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Aws.Ecs.Cluster("foo", new Aws.Ecs.ClusterArgs
        {
        });
    }

}

