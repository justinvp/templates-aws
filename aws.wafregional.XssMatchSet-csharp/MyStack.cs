using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var xssMatchSet = new Aws.WafRegional.XssMatchSet("xssMatchSet", new Aws.WafRegional.XssMatchSetArgs
        {
            XssMatchTuples = 
            {
                new Aws.WafRegional.Inputs.XssMatchSetXssMatchTupleArgs
                {
                    FieldToMatch = new Aws.WafRegional.Inputs.XssMatchSetXssMatchTupleFieldToMatchArgs
                    {
                        Type = "URI",
                    },
                    TextTransformation = "NONE",
                },
                new Aws.WafRegional.Inputs.XssMatchSetXssMatchTupleArgs
                {
                    FieldToMatch = new Aws.WafRegional.Inputs.XssMatchSetXssMatchTupleFieldToMatchArgs
                    {
                        Type = "QUERY_STRING",
                    },
                    TextTransformation = "NONE",
                },
            },
        });
    }

}

