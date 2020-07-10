import pulumi
import pulumi_aws as aws
import pulumi_pulumi as pulumi

alternate = pulumi.providers.Aws("alternate", profile="profile1")
sender_share = aws.ram.ResourceShare("senderShare",
    allow_external_principals=True,
    tags={
        "Name": "tf-test-resource-share",
    },
    opts=ResourceOptions(provider="aws.alternate"))
receiver = aws.get_caller_identity()
sender_invite = aws.ram.PrincipalAssociation("senderInvite",
    principal=receiver.account_id,
    resource_share_arn=sender_share.arn,
    opts=ResourceOptions(provider="aws.alternate"))
receiver_accept = aws.ram.ResourceShareAccepter("receiverAccept", share_arn=sender_invite.resource_share_arn)

