using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var sizeConstraintSet = new Aws.Waf.SizeConstraintSet("sizeConstraintSet", new Aws.Waf.SizeConstraintSetArgs
        {
            SizeConstraints = 
            {
                new Aws.Waf.Inputs.SizeConstraintSetSizeConstraintArgs
                {
                    ComparisonOperator = "EQ",
                    FieldToMatch = new Aws.Waf.Inputs.SizeConstraintSetSizeConstraintFieldToMatchArgs
                    {
                        Type = "BODY",
                    },
                    Size = 4096,
                    TextTransformation = "NONE",
                },
            },
        });
    }

}

