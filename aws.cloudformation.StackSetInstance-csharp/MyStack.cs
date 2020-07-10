using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.CloudFormation.StackSetInstance("example", new Aws.CloudFormation.StackSetInstanceArgs
        {
            AccountId = "123456789012",
            Region = "us-east-1",
            StackSetName = aws_cloudformation_stack_set.Example.Name,
        });
    }

}

