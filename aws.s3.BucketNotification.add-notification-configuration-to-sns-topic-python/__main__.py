import pulumi
import pulumi_aws as aws

bucket = aws.s3.Bucket("bucket")
topic = aws.sns.Topic("topic", policy=bucket.arn.apply(lambda arn: f"""{{
    "Version":"2012-10-17",
    "Statement":[{{
        "Effect": "Allow",
        "Principal": {{"AWS":"*"}},
        "Action": "SNS:Publish",
        "Resource": "arn:aws:sns:*:*:s3-event-notification-topic",
        "Condition":{{
            "ArnLike":{{"aws:SourceArn":"{arn}"}}
        }}
    }}]
}}

"""))
bucket_notification = aws.s3.BucketNotification("bucketNotification",
    bucket=bucket.id,
    topics=[{
        "events": ["s3:ObjectCreated:*"],
        "filterSuffix": ".log",
        "topic_arn": topic.arn,
    }])

