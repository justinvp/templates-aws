using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Waf.RegexPatternSet("example", new Aws.Waf.RegexPatternSetArgs
        {
            RegexPatternStrings = 
            {
                "one",
                "two",
            },
        });
    }

}

