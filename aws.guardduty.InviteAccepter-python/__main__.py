import pulumi
import pulumi_aws as aws

master = aws.guardduty.Detector("master")
member_detector = aws.guardduty.Detector("memberDetector", opts=ResourceOptions(provider="aws.dev"))
dev = aws.guardduty.Member("dev",
    account_id=member_detector.account_id,
    detector_id=master.id,
    email="required@example.com",
    invite=True)
member_invite_accepter = aws.guardduty.InviteAccepter("memberInviteAccepter",
    detector_id=member_detector.id,
    master_account_id=master.account_id,
    opts=ResourceOptions(provider="aws.dev",
        depends_on=["aws_guardduty_member.dev"]))

