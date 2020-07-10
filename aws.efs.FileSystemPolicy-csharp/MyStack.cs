using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var fs = new Aws.Efs.FileSystem("fs", new Aws.Efs.FileSystemArgs
        {
        });
        var policy = new Aws.Efs.FileSystemPolicy("policy", new Aws.Efs.FileSystemPolicyArgs
        {
            FileSystemId = fs.Id,
            Policy = @$"{{
    ""Version"": ""2012-10-17"",
    ""Id"": ""ExamplePolicy01"",
    ""Statement"": [
        {{
            ""Sid"": ""ExampleSatement01"",
            ""Effect"": ""Allow"",
            ""Principal"": {{
                ""AWS"": ""*""
            }},
            ""Resource"": ""{aws_efs_file_system.Test.Arn}"",
            ""Action"": [
                ""elasticfilesystem:ClientMount"",
                ""elasticfilesystem:ClientWrite""
            ],
            ""Condition"": {{
                ""Bool"": {{
                    ""aws:SecureTransport"": ""true""
                }}
            }}
        }}
    ]
}}

",
        });
    }

}

