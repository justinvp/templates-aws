import pulumi
import pulumi_aws as aws

test = aws.resourcegroups.Group("test", resource_query={
    "query": """{
  "ResourceTypeFilters": [
    "AWS::EC2::Instance"
  ],
  "TagFilters": [
    {
      "Key": "Stage",
      "Values": ["Test"]
    }
  ]
}

""",
})

