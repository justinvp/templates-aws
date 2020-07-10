using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bar = new Aws.Inspector.ResourceGroup("bar", new Aws.Inspector.ResourceGroupArgs
        {
            Tags = 
            {
                { "Env", "bar" },
                { "Name", "foo" },
            },
        });
        var foo = new Aws.Inspector.AssessmentTarget("foo", new Aws.Inspector.AssessmentTargetArgs
        {
            ResourceGroupArn = bar.Arn,
        });
    }

}

