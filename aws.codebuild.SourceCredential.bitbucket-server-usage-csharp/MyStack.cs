using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.CodeBuild.SourceCredential("example", new Aws.CodeBuild.SourceCredentialArgs
        {
            AuthType = "BASIC_AUTH",
            ServerType = "BITBUCKET",
            Token = "example",
            UserName = "test-user",
        });
    }

}

