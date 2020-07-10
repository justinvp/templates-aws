using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testTargetGroup = new Aws.LB.TargetGroup("testTargetGroup", new Aws.LB.TargetGroupArgs
        {
        });
        var testInstance = new Aws.Ec2.Instance("testInstance", new Aws.Ec2.InstanceArgs
        {
        });
        var testTargetGroupAttachment = new Aws.LB.TargetGroupAttachment("testTargetGroupAttachment", new Aws.LB.TargetGroupAttachmentArgs
        {
            Port = 80,
            TargetGroupArn = testTargetGroup.Arn,
            TargetId = testInstance.Id,
        });
    }

}

