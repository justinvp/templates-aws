import pulumi
import pulumi_aws as aws

queue = aws.sqs.Queue("queue")
test = aws.sqs.QueuePolicy("test",
    policy=queue.arn.apply(lambda arn: f"""{{
  "Version": "2012-10-17",
  "Id": "sqspolicy",
  "Statement": [
    {{
      "Sid": "First",
      "Effect": "Allow",
      "Principal": "*",
      "Action": "sqs:SendMessage",
      "Resource": "{arn}",
      "Condition": {{
        "ArnEquals": {{
          "aws:SourceArn": "{aws_sns_topic["example"]["arn"]}"
        }}
      }}
    }}
  ]
}}

"""),
    queue_url=queue.id)

