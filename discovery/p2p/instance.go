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
	"net"

	"github.com/cloudwego/kitex/pkg/discovery"
)

var (
	_ discovery.Instance = &Instance{}
	_ net.Addr           = &NetAddr{}
)

type Instance struct {
	netAddr net.Addr
	weight  int
}

func NewInstance(netAddr net.Addr, weight int) discovery.Instance {
	return &Instance{
		netAddr: netAddr,
		weight:  weight,
	}
}

func (i *Instance) Address() net.Addr {
	return i.netAddr
}

func (i *Instance) Weight() int {
	return i.weight
}

func (i *Instance) Tag(key string) (value string, exist bool) {
	return "", false
}

type NetAddr struct {
	network string
	address string
}

func NewNetAddr(network, address string) net.Addr {
	return &NetAddr{network, address}
}

func (na *NetAddr) Network() string {
	return na.network
}

func (na *NetAddr) String() string {
	return na.address
}
