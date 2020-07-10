import pulumi
import pulumi_aws as aws

main = aws.elb.get_service_account()
elb_logs = aws.s3.Bucket("elbLogs",
    acl="private",
    policy=f"""{{
  "Id": "Policy",
  "Version": "2012-10-17",
  "Statement": [
    {{
      "Action": [
        "s3:PutObject"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::my-elb-tf-test-bucket/AWSLogs/*",
      "Principal": {{
        "AWS": [
          "{main.arn}"
        ]
      }}
    }}
  ]
}}

""")
bar = aws.elb.LoadBalancer("bar",
    access_logs={
        "bucket": elb_logs.bucket,
        "interval": 5,
    },
    availability_zones=["us-west-2a"],
    listeners=[{
        "instance_port": 8000,
        "instanceProtocol": "http",
        "lb_port": 80,
        "lbProtocol": "http",
    }])

