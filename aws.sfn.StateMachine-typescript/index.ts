import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sfnStateMachine = new aws.sfn.StateMachine("sfn_state_machine", {
    definition: pulumi.interpolate`{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
  "StartAt": "HelloWorld",
  "States": {
    "HelloWorld": {
      "Type": "Task",
      "Resource": "${aws_lambda_function_lambda.arn}",
      "End": true
    }
  }
}
`,
    roleArn: aws_iam_role_iam_for_sfn.arn,
});

