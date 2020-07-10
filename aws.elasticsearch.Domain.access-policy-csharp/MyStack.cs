using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();
        var domain = config.Get("domain") ?? "tf-test";
        var currentRegion = Output.Create(Aws.GetRegion.InvokeAsync());
        var currentCallerIdentity = Output.Create(Aws.GetCallerIdentity.InvokeAsync());
        var example = new Aws.ElasticSearch.Domain("example", new Aws.ElasticSearch.DomainArgs
        {
            AccessPolicies = Output.Tuple(currentRegion, currentCallerIdentity).Apply(values =>
            {
                var currentRegion = values.Item1;
                var currentCallerIdentity = values.Item2;
                return @$"{{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {{
      ""Action"": ""es:*"",
      ""Principal"": ""*"",
      ""Effect"": ""Allow"",
      ""Resource"": ""arn:aws:es:{currentRegion.Name}:{currentCallerIdentity.AccountId}:domain/{domain}/*"",
      ""Condition"": {{
        ""IpAddress"": {{""aws:SourceIp"": [""66.193.100.22/32""]}}
      }}
    }}
  ]
}}

";
            }),
        });
    }

}

