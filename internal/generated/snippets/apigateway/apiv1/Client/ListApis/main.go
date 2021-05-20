// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START apigateway_v1_generated_ApiGatewayService_ListApis_sync]

package main

import (
	"context"

	apigateway "cloud.google.com/go/apigateway/apiv1"
	"google.golang.org/api/iterator"
	apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"
)

func main() {
	ctx := context.Background()
	c, err := apigateway.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &apigatewaypb.ListApisRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListApis(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

// [END apigateway_v1_generated_ApiGatewayService_ListApis_sync]