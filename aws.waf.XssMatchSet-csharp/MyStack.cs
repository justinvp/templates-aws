using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var xssMatchSet = new Aws.Waf.XssMatchSet("xssMatchSet", new Aws.Waf.XssMatchSetArgs
        {
            XssMatchTuples = 
            {
                new Aws.Waf.Inputs.XssMatchSetXssMatchTupleArgs
                {
                    FieldToMatch = new Aws.Waf.Inputs.XssMatchSetXssMatchTupleFieldToMatchArgs
                    {
                        Type = "URI",
                    },
                    TextTransformation = "NONE",
                },
                new Aws.Waf.Inputs.XssMatchSetXssMatchTupleArgs
                {
                    FieldToMatch = new Aws.Waf.Inputs.XssMatchSetXssMatchTupleFieldToMatchArgs
                    {
                        Type = "QUERY_STRING",
                    },
                    TextTransformation = "NONE",
                },
            },
        });
    }

}

