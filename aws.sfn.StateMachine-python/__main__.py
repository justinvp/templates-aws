import pulumi
import pulumi_aws as aws

sfn_state_machine = aws.sfn.StateMachine("sfnStateMachine",
    definition=f"""{{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
  "StartAt": "HelloWorld",
  "States": {{
    "HelloWorld": {{
      "Type": "Task",
      "Resource": "{aws_lambda_function["lambda"]["arn"]}",
      "End": true
    }}
  }}
}}

""",
    role_arn=aws_iam_role["iam_for_sfn"]["arn"])

