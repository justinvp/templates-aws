import pulumi
import pulumi_aws as aws

example = aws.macie.MemberAccountAssociation("example", member_account_id="123456789012")

