using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var master = new Aws.GuardDuty.Detector("master", new Aws.GuardDuty.DetectorArgs
        {
        });
        var memberDetector = new Aws.GuardDuty.Detector("memberDetector", new Aws.GuardDuty.DetectorArgs
        {
        }, new CustomResourceOptions
        {
            Provider = "aws.dev",
        });
        var dev = new Aws.GuardDuty.Member("dev", new Aws.GuardDuty.MemberArgs
        {
            AccountId = memberDetector.AccountId,
            DetectorId = master.Id,
            Email = "required@example.com",
            Invite = true,
        });
        var memberInviteAccepter = new Aws.GuardDuty.InviteAccepter("memberInviteAccepter", new Aws.GuardDuty.InviteAccepterArgs
        {
            DetectorId = memberDetector.Id,
            MasterAccountId = master.AccountId,
        }, new CustomResourceOptions
        {
            Provider = "aws.dev",
            DependsOn = 
            {
                "aws_guardduty_member.dev",
            },
        });
    }

}

