import pulumi
import pulumi_aws as aws

my_detector = aws.guardduty.Detector("myDetector", enable=True)

