using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        // Create a new load balancer attachment
        var asgAttachmentBar = new Aws.AutoScaling.Attachment("asgAttachmentBar", new Aws.AutoScaling.AttachmentArgs
        {
            AutoscalingGroupName = aws_autoscaling_group.Asg.Id,
            Elb = aws_elb.Bar.Id,
        });
    }

}

