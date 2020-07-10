using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var fooWithLifecylePolicy = new Aws.Efs.FileSystem("fooWithLifecylePolicy", new Aws.Efs.FileSystemArgs
        {
            LifecyclePolicy = new Aws.Efs.Inputs.FileSystemLifecyclePolicyArgs
            {
                TransitionToIa = "AFTER_30_DAYS",
            },
        });
    }

}

