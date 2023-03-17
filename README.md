# bootstrap-terraform-s3-backend

An AWS CDK tool used to bootstrap an S3 Terraform backend with DynamoDB for state locking.

## Environment Variables

Functionality can be configured by setting these environment variables before synthesizing the Cloudformation stack definition.

| Variable Name              |
| -------------------------- |
| CDK_AWS_ACCOUNT_ID         |
| CDK_AWS_REGION             |
| CDK_TF_BACKEND_BUCKET_NAME |
| CDK_TF_LOCK_TABLE_NAME     |

## Useful commands

* `cdk deploy`      deploy this stack to your default AWS account/region
* `cdk diff`        compare deployed stack with current state
* `cdk synth`       emits the synthesized CloudFormation template
* `go test`         run unit tests
