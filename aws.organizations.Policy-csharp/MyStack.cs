using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Organizations.Policy("example", new Aws.Organizations.PolicyArgs
        {
            Content = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": {
    ""Effect"": ""Allow"",
    ""Action"": ""*"",
    ""Resource"": ""*""
  }
}

",
        });
    }

}

