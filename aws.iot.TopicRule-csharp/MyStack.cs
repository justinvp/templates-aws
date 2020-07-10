using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var mytopic = new Aws.Sns.Topic("mytopic", new Aws.Sns.TopicArgs
        {
        });
        var role = new Aws.Iam.Role("role", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Effect"": ""Allow"",
      ""Principal"": {
        ""Service"": ""iot.amazonaws.com""
      },
      ""Action"": ""sts:AssumeRole""
    }
  ]
}

",
        });
        var rule = new Aws.Iot.TopicRule("rule", new Aws.Iot.TopicRuleArgs
        {
            Description = "Example rule",
            Enabled = true,
            Sns = new Aws.Iot.Inputs.TopicRuleSnsArgs
            {
                Sns = "RAW",
                Sns = role.Arn,
                Sns = mytopic.Arn,
            },
            Sql = "SELECT * FROM 'topic/test'",
            SqlVersion = "2016-03-23",
        });
        var iamPolicyForLambda = new Aws.Iam.RolePolicy("iamPolicyForLambda", new Aws.Iam.RolePolicyArgs
        {
            Policy = mytopic.Arn.Apply(arn => @$"{{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {{
        ""Effect"": ""Allow"",
        ""Action"": [
            ""sns:Publish""
        ],
        ""Resource"": ""{arn}""
    }}
  ]
}}

"),
            Role = role.Id,
        });
    }

}

