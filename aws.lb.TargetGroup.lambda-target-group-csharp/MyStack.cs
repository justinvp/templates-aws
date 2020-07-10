using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var lambda_example = new Aws.LB.TargetGroup("lambda-example", new Aws.LB.TargetGroupArgs
        {
            TargetType = "lambda",
        });
    }

}

