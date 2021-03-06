module github.com/bigkraig/pulumi-infoblox

go 1.12

replace (
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
)

require (
	github.com/hashicorp/terraform v0.12.0-alpha4.0.20190401213546-16778fea9219
	github.com/pkg/errors v0.8.0
	github.com/pulumi/pulumi v0.17.6-0.20190410045519-ef5e148a73c0
	github.com/pulumi/pulumi-terraform v0.14.1-dev.0.20190410175858-788550dffb09
	github.com/sky-uk/go-rest-api v0.0.0-20181002093433-b85c2e6be634 // indirect
	github.com/sky-uk/skyinfoblox v0.0.0-20171219155614-771f916148fa // indirect
	github.com/sky-uk/terraform-provider-infoblox v0.0.0-20171101151550-83f446ed9e90
	github.com/stretchr/testify v1.3.0
)
