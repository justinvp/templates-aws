package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/datasync"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datasync.NewEfsLocation(ctx, "example", &datasync.EfsLocationArgs{
			Ec2Config: &datasync.EfsLocationEc2ConfigArgs{
				SecurityGroupArns: pulumi.StringArray{
					pulumi.Any(aws_security_group.Example.Arn),
				},
				SubnetArn: pulumi.Any(aws_subnet.Example.Arn),
			},
			EfsFileSystemArn: pulumi.Any(aws_efs_mount_target.Example.File_system_arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

