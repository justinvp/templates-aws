using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var sfnActivity = new Aws.Sfn.Activity("sfnActivity", new Aws.Sfn.ActivityArgs
        {
        });
    }

}

