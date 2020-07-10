using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bar = new Aws.Dax.Cluster("bar", new Aws.Dax.ClusterArgs
        {
            ClusterName = "cluster-example",
            IamRoleArn = data.Aws_iam_role.Example.Arn,
            NodeType = "dax.r4.large",
            ReplicationFactor = 1,
        });
    }

}

