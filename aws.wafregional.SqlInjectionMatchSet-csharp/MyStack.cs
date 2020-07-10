using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var sqlInjectionMatchSet = new Aws.WafRegional.SqlInjectionMatchSet("sqlInjectionMatchSet", new Aws.WafRegional.SqlInjectionMatchSetArgs
        {
            SqlInjectionMatchTuples = 
            {
                new Aws.WafRegional.Inputs.SqlInjectionMatchSetSqlInjectionMatchTupleArgs
                {
                    FieldToMatch = new Aws.WafRegional.Inputs.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatchArgs
                    {
                        Type = "QUERY_STRING",
                    },
                    TextTransformation = "URL_DECODE",
                },
            },
        });
    }

}

