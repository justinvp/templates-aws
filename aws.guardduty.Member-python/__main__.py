import pulumi
import pulumi_aws as aws

master = aws.guardduty.Detector("master", enable=True)
member_detector = aws.guardduty.Detector("memberDetector", enable=True,
opts=ResourceOptions(provider="aws.dev"))
member_member = aws.guardduty.Member("memberMember",
    account_id=member_detector.account_id,
    detector_id=master.id,
    email="required@example.com",
    invite=True,
    invitation_message="please accept guardduty invitation")

