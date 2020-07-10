using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var master = new Aws.GuardDuty.Detector("master", new Aws.GuardDuty.DetectorArgs
        {
            Enable = true,
        });
        var memberDetector = new Aws.GuardDuty.Detector("memberDetector", new Aws.GuardDuty.DetectorArgs
        {
            Enable = true,
        }, new CustomResourceOptions
        {
            Provider = "aws.dev",
        });
        var memberMember = new Aws.GuardDuty.Member("memberMember", new Aws.GuardDuty.MemberArgs
        {
            AccountId = memberDetector.AccountId,
            DetectorId = master.Id,
            Email = "required@example.com",
            Invite = true,
            InvitationMessage = "please accept guardduty invitation",
        });
    }

}

