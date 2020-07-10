using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var myDetector = new Aws.GuardDuty.Detector("myDetector", new Aws.GuardDuty.DetectorArgs
        {
            Enable = true,
        });
    }

}

