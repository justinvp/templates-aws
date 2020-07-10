import pulumi
import pulumi_aws as aws

app = aws.pinpoint.App("app")
sms = aws.pinpoint.SmsChannel("sms", application_id=app.application_id)

