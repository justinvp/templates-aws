using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleAccount = new Aws.SecurityHub.Account("exampleAccount", new Aws.SecurityHub.AccountArgs
        {
        });
        var exampleMember = new Aws.SecurityHub.Member("exampleMember", new Aws.SecurityHub.MemberArgs
        {
            AccountId = "123456789012",
            Email = "example@example.com",
            Invite = true,
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                "aws_securityhub_account.example",
            },
        });
    }

}

