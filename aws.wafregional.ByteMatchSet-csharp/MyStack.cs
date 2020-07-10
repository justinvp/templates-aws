using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var byteSet = new Aws.WafRegional.ByteMatchSet("byteSet", new Aws.WafRegional.ByteMatchSetArgs
        {
            ByteMatchTuples = 
            {
                new Aws.WafRegional.Inputs.ByteMatchSetByteMatchTupleArgs
                {
                    FieldToMatch = new Aws.WafRegional.Inputs.ByteMatchSetByteMatchTupleFieldToMatchArgs
                    {
                        Data = "referer",
                        Type = "HEADER",
                    },
                    PositionalConstraint = "CONTAINS",
                    TargetString = "badrefer1",
                    TextTransformation = "NONE",
                },
            },
        });
    }

}

