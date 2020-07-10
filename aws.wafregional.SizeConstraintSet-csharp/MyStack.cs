using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var sizeConstraintSet = new Aws.WafRegional.SizeConstraintSet("sizeConstraintSet", new Aws.WafRegional.SizeConstraintSetArgs
        {
            SizeConstraints = 
            {
                new Aws.WafRegional.Inputs.SizeConstraintSetSizeConstraintArgs
                {
                    ComparisonOperator = "EQ",
                    FieldToMatch = new Aws.WafRegional.Inputs.SizeConstraintSetSizeConstraintFieldToMatchArgs
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

