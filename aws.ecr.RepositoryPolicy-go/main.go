package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ecr"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		foo, err := ecr.NewRepository(ctx, "foo", nil)
		if err != nil {
			return err
		}
		_, err = ecr.NewRepositoryPolicy(ctx, "foopolicy", &ecr.RepositoryPolicyArgs{
			Policy:     pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\": \"2008-10-17\",\n", "    \"Statement\": [\n", "        {\n", "            \"Sid\": \"new policy\",\n", "            \"Effect\": \"Allow\",\n", "            \"Principal\": \"*\",\n", "            \"Action\": [\n", "                \"ecr:GetDownloadUrlForLayer\",\n", "                \"ecr:BatchGetImage\",\n", "                \"ecr:BatchCheckLayerAvailability\",\n", "                \"ecr:PutImage\",\n", "                \"ecr:InitiateLayerUpload\",\n", "                \"ecr:UploadLayerPart\",\n", "                \"ecr:CompleteLayerUpload\",\n", "                \"ecr:DescribeRepositories\",\n", "                \"ecr:GetRepositoryPolicy\",\n", "                \"ecr:ListImages\",\n", "                \"ecr:DeleteRepository\",\n", "                \"ecr:BatchDeleteImage\",\n", "                \"ecr:SetRepositoryPolicy\",\n", "                \"ecr:DeleteRepositoryPolicy\"\n", "            ]\n", "        }\n", "    ]\n", "}\n", "\n")),
			Repository: foo.Name,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

