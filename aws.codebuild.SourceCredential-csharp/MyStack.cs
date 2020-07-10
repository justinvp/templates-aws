using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.CodeBuild.SourceCredential("example", new Aws.CodeBuild.SourceCredentialArgs
        {
            AuthType = "PERSONAL_ACCESS_TOKEN",
            ServerType = "GITHUB",
            Token = "example",
        });
    }

}

