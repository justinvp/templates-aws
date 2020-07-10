using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ebs.DefaultKmsKey("example", new Aws.Ebs.DefaultKmsKeyArgs
        {
            KeyArn = aws_kms_key.Example.Arn,
        });
    }

}

