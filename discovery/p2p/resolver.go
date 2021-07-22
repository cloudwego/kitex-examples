// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package p2p

import (
	"context"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var _ discovery.Resolver = &P2PResolver{}

type P2PResolver struct {
	network string
	target  string
}

func NewP2PResolver(network, target string) discovery.Resolver {
	return &P2PResolver{network: network, target: target}
}

func (p *P2PResolver) Resolve(ctx context.Context, desc string) (discovery.Result, error) {
	instances := []discovery.Instance{NewInstance(NewNetAddr(p.network, p.target), 100)}
	return discovery.Result{Cacheable: true, CacheKey: desc, Instances: instances}, nil
}

// Target should return a description for the given target that is suitable for being a key for cache.
func (p *P2PResolver) Target(ctx context.Context, target rpcinfo.EndpointInfo) (description string) {
	return p.network + "@" + p.target
}

// Name returns the name of the resolver.
func (p *P2PResolver) Name() string {
	return "P2PResolver"
}

// Diff computes the difference between two results.
// When `next` is cacheable, the Change should be cacheable, too. And the `Result` field's CacheKey in
// the return value should be set with the given cacheKey.
func (p *P2PResolver) Diff(cacheKey string, prev discovery.Result, next discovery.Result) (discovery.Change, bool) {
	return discovery.DefaultDiff(cacheKey, prev, next)
}
