using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ElasticSearch.Domain("example", new Aws.ElasticSearch.DomainArgs
        {
            ElasticsearchVersion = "2.3",
        });
        var main = new Aws.ElasticSearch.DomainPolicy("main", new Aws.ElasticSearch.DomainPolicyArgs
        {
            AccessPolicies = example.Arn.Apply(arn => @$"{{
    ""Version"": ""2012-10-17"",
    ""Statement"": [
        {{
            ""Action"": ""es:*"",
            ""Principal"": ""*"",
            ""Effect"": ""Allow"",
            ""Condition"": {{
                ""IpAddress"": {{""aws:SourceIp"": ""127.0.0.1/32""}}
            }},
            ""Resource"": ""{arn}/*""
        }}
    ]
}}

"),
            DomainName = example.DomainName,
        });
    }

}

