import pulumi
import pulumi_aws as aws

bucket = aws.s3.Bucket("bucket")
queue = aws.sqs.Queue("queue", policy=bucket.arn.apply(lambda arn: f"""{{
  "Version": "2012-10-17",
  "Statement": [
    {{
      "Effect": "Allow",
      "Principal": "*",
      "Action": "sqs:SendMessage",
	  "Resource": "arn:aws:sqs:*:*:s3-event-notification-queue",
      "Condition": {{
        "ArnEquals": {{ "aws:SourceArn": "{arn}" }}
      }}
    }}
  ]
}}

"""))
bucket_notification = aws.s3.BucketNotification("bucketNotification",
    bucket=bucket.id,
    queues=[
        {
            "events": ["s3:ObjectCreated:*"],
            "filterPrefix": "images/",
            "id": "image-upload-event",
            "queueArn": queue.arn,
        },
        {
            "events": ["s3:ObjectCreated:*"],
            "filterPrefix": "videos/",
            "id": "video-upload-event",
            "queueArn": queue.arn,
        },
    ])

