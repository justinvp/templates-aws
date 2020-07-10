using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var my_instance = new Aws.OpsWorks.Instance("my-instance", new Aws.OpsWorks.InstanceArgs
        {
            InstanceType = "t2.micro",
            LayerIds = 
            {
                aws_opsworks_custom_layer.My_layer.Id,
            },
            Os = "Amazon Linux 2015.09",
            StackId = aws_opsworks_stack.Main.Id,
            State = "stopped",
        });
    }

}

