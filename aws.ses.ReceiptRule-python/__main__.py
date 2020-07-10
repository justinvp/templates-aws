import pulumi
import pulumi_aws as aws

# Add a header to the email and store it in S3
store = aws.ses.ReceiptRule("store",
    add_header_actions=[{
        "headerName": "Custom-Header",
        "headerValue": "Added by SES",
        "position": 1,
    }],
    enabled=True,
    recipients=["karen@example.com"],
    rule_set_name="default-rule-set",
    s3_actions=[{
        "bucket_name": "emails",
        "position": 2,
    }],
    scan_enabled=True)

