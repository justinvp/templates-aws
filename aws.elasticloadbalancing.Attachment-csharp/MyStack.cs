using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        // Create a new load balancer attachment
        var baz = new Aws.Elb.Attachment("baz", new Aws.Elb.AttachmentArgs
        {
            Elb = aws_elb.Bar.Id,
            Instance = aws_instance.Foo.Id,
        });
    }

}

