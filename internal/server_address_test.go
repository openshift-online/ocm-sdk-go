/*
Copyright (c) 2021 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package internal

import (
	"net/url"

	. "github.com/onsi/ginkgo/v2/dsl/table" // nolint
	. "github.com/onsi/gomega"              // nolint
)

var _ = DescribeTable(
	"Parsing success",
	func(input string, expected *ServerAddress) {
		actual, err := ParseServerAddress(input)
		Expect(err).ToNot(HaveOccurred())
		Expect(actual.Text).To(Equal(input))
		Expect(actual.Network).To(Equal(expected.Network))
		Expect(actual.Protocol).To(Equal(expected.Protocol))
		Expect(actual.Host).To(Equal(expected.Host))
		Expect(actual.Port).To(Equal(expected.Port))
		Expect(actual.Socket).To(Equal(expected.Socket))
		Expect(actual.URL.String()).To(Equal(expected.URL.String()))
	},
	Entry(
		"tcp",
		"tcp://my.server.com",
		&ServerAddress{
			Network:  TCPNetwork,
			Protocol: HTTPProtocol,
			Host:     "my.server.com",
			Port:     "80",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com",
			},
		},
	),
	Entry(
		"tcp+http",
		"tcp+http://my.server.com",
		&ServerAddress{
			Network:  TCPNetwork,
			Protocol: HTTPProtocol,
			Host:     "my.server.com",
			Port:     "80",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com",
			},
		},
	),
	Entry(
		"tcp+https",
		"tcp+https://my.server.com",
		&ServerAddress{
			Network:  TCPNetwork,
			Protocol: HTTPSProtocol,
			Host:     "my.server.com",
			Port:     "443",
			URL: &url.URL{
				Scheme: "https",
				Host:   "my.server.com",
			},
		},
	),
	Entry(
		"tcp+h2c",
		"tcp+h2c://my.server.com",
		&ServerAddress{
			Network:  TCPNetwork,
			Protocol: H2CProtocol,
			Host:     "my.server.com",
			Port:     "80",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com",
			},
		},
	),
	Entry(
		"unix",
		"unix://my.server.com/my.socket",
		&ServerAddress{
			Network:  UnixNetwork,
			Protocol: HTTPProtocol,
			Host:     "my.server.com",
			Port:     "80",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com",
			},
			Socket: "/my.socket",
		},
	),
	Entry(
		"unix+http",
		"unix+http://my.server.com/my.socket",
		&ServerAddress{
			Network:  UnixNetwork,
			Protocol: HTTPProtocol,
			Host:     "my.server.com",
			Port:     "80",
			Socket:   "/my.socket",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com",
			},
		},
	),
	Entry(
		"unix+https",
		"unix+https://my.server.com/my.socket",
		&ServerAddress{
			Network:  UnixNetwork,
			Protocol: HTTPSProtocol,
			Host:     "my.server.com",
			Port:     "443",
			Socket:   "/my.socket",
			URL: &url.URL{
				Scheme: "https",
				Host:   "my.server.com",
			},
		},
	),
	Entry(
		"unix+h2c",
		"unix+h2c://my.server.com/my.socket",
		&ServerAddress{
			Network:  UnixNetwork,
			Protocol: H2CProtocol,
			Host:     "my.server.com",
			Port:     "80",
			Socket:   "/my.socket",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com",
			},
		},
	),
	Entry(
		"http",
		"http://my.server.com",
		&ServerAddress{
			Network:  TCPNetwork,
			Protocol: HTTPProtocol,
			Host:     "my.server.com",
			Port:     "80",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com",
			},
		},
	),
	Entry(
		"http+tcp",
		"http+tcp://my.server.com",
		&ServerAddress{
			Network:  TCPNetwork,
			Protocol: HTTPProtocol,
			Host:     "my.server.com",
			Port:     "80",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com",
			},
		},
	),
	Entry(
		"http+unix",
		"http+unix://my.server.com/my.socket",
		&ServerAddress{
			Network:  UnixNetwork,
			Protocol: HTTPProtocol,
			Host:     "my.server.com",
			Port:     "80",
			Socket:   "/my.socket",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com",
				Path:   "",
			},
		},
	),
	Entry(
		"https",
		"https://my.server.com",
		&ServerAddress{
			Network:  TCPNetwork,
			Protocol: HTTPSProtocol,
			Host:     "my.server.com",
			Port:     "443",
			URL: &url.URL{
				Scheme: "https",
				Host:   "my.server.com",
			},
		},
	),
	Entry(
		"https+tcp",
		"https+tcp://my.server.com",
		&ServerAddress{
			Network:  TCPNetwork,
			Protocol: HTTPSProtocol,
			Host:     "my.server.com",
			Port:     "443",
			URL: &url.URL{
				Scheme: "https",
				Host:   "my.server.com",
			},
		},
	),
	Entry(
		"https+unix",
		"https+unix://my.server.com/my.socket",
		&ServerAddress{
			Network:  UnixNetwork,
			Protocol: HTTPSProtocol,
			Host:     "my.server.com",
			Port:     "443",
			Socket:   "/my.socket",
			URL: &url.URL{
				Scheme: "https",
				Host:   "my.server.com",
			},
		},
	),
	Entry(
		"h2c",
		"h2c://my.server.com",
		&ServerAddress{
			Network:  TCPNetwork,
			Protocol: H2CProtocol,
			Host:     "my.server.com",
			Port:     "80",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com",
			},
		},
	),
	Entry(
		"h2c+tcp",
		"h2c+tcp://my.server.com",
		&ServerAddress{
			Network:  TCPNetwork,
			Protocol: H2CProtocol,
			Host:     "my.server.com",
			Port:     "80",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com",
			},
		},
	),
	Entry(
		"h2c+unix",
		"h2c+unix://my.server.com/my.socket",
		&ServerAddress{
			Network:  UnixNetwork,
			Protocol: H2CProtocol,
			Host:     "my.server.com",
			Port:     "80",
			Socket:   "/my.socket",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com",
			},
		},
	),
	Entry(
		"Non default HTTP port",
		"http://my.server.com:1080",
		&ServerAddress{
			Network:  TCPNetwork,
			Protocol: HTTPProtocol,
			Host:     "my.server.com",
			Port:     "1080",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com:1080",
			},
		},
	),
	Entry(
		"Non default HTTPS port",
		"http://my.server.com:1443",
		&ServerAddress{
			Network:  TCPNetwork,
			Protocol: HTTPProtocol,
			Host:     "my.server.com",
			Port:     "1443",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com:1443",
			},
		},
	),
	Entry(
		"Non default H2C port",
		"h2c://my.server.com:1080",
		&ServerAddress{
			Network:  TCPNetwork,
			Protocol: H2CProtocol,
			Host:     "my.server.com",
			Port:     "1080",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com:1080",
			},
		},
	),
	Entry(
		"Unix socket in query parameter",
		"unix://my.server.com/my/path?socket=/my.socket",
		&ServerAddress{
			Network:  UnixNetwork,
			Protocol: HTTPProtocol,
			Host:     "my.server.com",
			Port:     "80",
			Socket:   "/my.socket",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com",
			},
		},
	),
	Entry(
		"TCP network from query parameter",
		"http://my.server.com?network=tcp",
		&ServerAddress{
			Network:  TCPNetwork,
			Protocol: HTTPProtocol,
			Host:     "my.server.com",
			Port:     "80",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com",
			},
		},
	),
	Entry(
		"Unix network from query parameter",
		"http://my.server.com/my/path?network=unix&socket=/my.socket",
		&ServerAddress{
			Network:  UnixNetwork,
			Protocol: HTTPProtocol,
			Host:     "my.server.com",
			Port:     "80",
			Socket:   "/my.socket",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com",
			},
		},
	),
	Entry(
		"H2C protocol from query parameter",
		"tcp://my.server.com?protocol=h2c",
		&ServerAddress{
			Network:  TCPNetwork,
			Protocol: H2CProtocol,
			Host:     "my.server.com",
			Port:     "80",
			URL: &url.URL{
				Scheme: "http",
				Host:   "my.server.com",
			},
		},
	),
)

var _ = DescribeTable(
	"Parsing error",
	func(input string, expected ...string) {
		actual, err := ParseServerAddress(input)
		Expect(err).To(HaveOccurred())
		Expect(actual).To(BeNil())
		message := err.Error()
		for _, substring := range expected {
			Expect(message).To(ContainSubstring(substring))
		}
	},
	Entry(
		"Unknonwn network",
		"mynet+http://my.server.com",
		"component 'mynet' of scheme 'mynet+http' doesn't correspond to any supported "+
			"network or protocol",
		"supported networks are 'tcp' and 'unix'",
	),
	Entry(
		"Unknonwn protocol",
		"tcp+myprotocol://my.server.com",
		"component 'myprotocol' of scheme 'tcp+myprotocol' doesn't correspond to any "+
			"supported network or protocol",
		"supported protocols are 'http', 'https' and 'h2c'",
	),
	Entry(
		"Missing Unix socket",
		"unix://my.server.com",
		"expected socket name in the 'socket' query parameter or in the path but both "+
			"are empty",
	),
	Entry(
		"Incompatible network from query parameter",
		"unix://my.server.com/my.socket?network=tcp",
		"network 'tcp' from query parameter isn't compatible with network 'unix' "+
			"from scheme",
	),
	Entry(
		"Invalid network from query parameter",
		"http://my.server.com/my.socket?network=unox",
		"network 'unox' isn't valid, valid values are 'unix' and 'tcp'",
	),
	Entry(
		"Invalid protocol from query parameter",
		"tcp://my.server.com/my.socket?protocol=h2d",
		"protocol 'h2d' isn't valid, valid values are 'http', 'https' and 'h2c'",
	),
)
