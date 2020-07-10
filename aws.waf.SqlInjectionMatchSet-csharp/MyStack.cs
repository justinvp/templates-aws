using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var sqlInjectionMatchSet = new Aws.Waf.SqlInjectionMatchSet("sqlInjectionMatchSet", new Aws.Waf.SqlInjectionMatchSetArgs
        {
            SqlInjectionMatchTuples = 
            {
                new Aws.Waf.Inputs.SqlInjectionMatchSetSqlInjectionMatchTupleArgs
                {
                    FieldToMatch = new Aws.Waf.Inputs.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatchArgs
                    {
                        Type = "QUERY_STRING",
                    },
                    TextTransformation = "URL_DECODE",
                },
            },
        });
    }

}

