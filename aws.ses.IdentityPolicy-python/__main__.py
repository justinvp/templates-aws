import pulumi
import pulumi_aws as aws

example_domain_identity = aws.ses.DomainIdentity("exampleDomainIdentity", domain="example.com")
example_policy_document = example_domain_identity.arn.apply(lambda arn: aws.iam.get_policy_document(statements=[{
    "actions": [
        "SES:SendEmail",
        "SES:SendRawEmail",
    ],
    "principals": [{
        "identifiers": ["*"],
        "type": "AWS",
    }],
    "resources": [arn],
}]))
example_identity_policy = aws.ses.IdentityPolicy("exampleIdentityPolicy",
    identity=example_domain_identity.arn,
    policy=example_policy_document.json)

