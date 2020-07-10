using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foo = new Aws.Ecr.Repository("foo", new Aws.Ecr.RepositoryArgs
        {
        });
        var foopolicy = new Aws.Ecr.RepositoryPolicy("foopolicy", new Aws.Ecr.RepositoryPolicyArgs
        {
            Policy = @"{
    ""Version"": ""2008-10-17"",
    ""Statement"": [
        {
            ""Sid"": ""new policy"",
            ""Effect"": ""Allow"",
            ""Principal"": ""*"",
            ""Action"": [
                ""ecr:GetDownloadUrlForLayer"",
                ""ecr:BatchGetImage"",
                ""ecr:BatchCheckLayerAvailability"",
                ""ecr:PutImage"",
                ""ecr:InitiateLayerUpload"",
                ""ecr:UploadLayerPart"",
                ""ecr:CompleteLayerUpload"",
                ""ecr:DescribeRepositories"",
                ""ecr:GetRepositoryPolicy"",
                ""ecr:ListImages"",
                ""ecr:DeleteRepository"",
                ""ecr:BatchDeleteImage"",
                ""ecr:SetRepositoryPolicy"",
                ""ecr:DeleteRepositoryPolicy""
            ]
        }
    ]
}

",
            Repository = foo.Name,
        });
    }

}

