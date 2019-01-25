module whd-grpc-demo

require (
	github.com/go-xorm/xorm v0.7.1
	github.com/golang/protobuf v1.2.0
	github.com/google/go-cmp v0.2.0 // indirect
	github.com/xprst/whd-grpc-base v0.0.0-20190124061501-b827118b64d6
	golang.org/x/crypto v0.0.0-20181203042331-505ab145d0a9 // indirect
	golang.org/x/net v0.0.0-20190110200230-915654e7eabc
	golang.org/x/sys v0.0.0-20181205085412-a5c9d58dba9a // indirect
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc v1.16.0
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.34.0
	go.uber.org/zap => github.com/uber-go/zap v1.9.1
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190122013713-64072686203f
	golang.org/x/lint => github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3
	golang.org/x/net => github.com/golang/net v0.0.0-20190110200230-915654e7eabc
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/sync => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20180830151530-49385e6e1522
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20180828015842-6cd1fcedba52
	google.golang.org/appengine => github.com/golang/appengine v1.1.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190111180523-db91494dd46c
	google.golang.org/grpc => github.com/grpc/grpc-go v1.2.1-0.20190114234020-98a94b0cb0eb
)
