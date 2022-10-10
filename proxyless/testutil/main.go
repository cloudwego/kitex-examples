package main

import "github.com/cloudwego/kitex-proxyless-test/testutil/testSuite"

func main() {
	tc := testSuite.NewController()
	tc.Run()
}
