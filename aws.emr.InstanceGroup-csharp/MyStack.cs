using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var task = new Aws.Emr.InstanceGroup("task", new Aws.Emr.InstanceGroupArgs
        {
            ClusterId = aws_emr_cluster.Tf_test_cluster.Id,
            InstanceCount = 1,
            InstanceType = "m5.xlarge",
        });
    }

}

