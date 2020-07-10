using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var byteSet = new Aws.Waf.ByteMatchSet("byteSet", new Aws.Waf.ByteMatchSetArgs
        {
            ByteMatchTuples = 
            {
                new Aws.Waf.Inputs.ByteMatchSetByteMatchTupleArgs
                {
                    FieldToMatch = new Aws.Waf.Inputs.ByteMatchSetByteMatchTupleFieldToMatchArgs
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

