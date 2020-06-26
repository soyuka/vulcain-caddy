// Copyright 2015 Matthew Holt and The Caddy Authors
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

package reverseproxy

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp/reverseproxy"
	// "github.com/dunglas/vulcain/gateway"
)

func init() {
	caddy.RegisterModule(Vulcain{})
}

type Vulcain struct {
	*reverseproxy.Handler
	// gateway *gateway.Gateway
}

// CaddyModule returns the Caddy module information.
func (Vulcain) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.vulcain",
		New: func() caddy.Module { 
			v := new(Vulcain)
			v.Handler = new(reverseproxy.Handler)
			return v
		},
	}
}

func (v *Vulcain) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	
	return v.Handler.ServeHTTP(w, r, next)
}

// Provision sets up h's configuration.
func (v *Vulcain) Provision(ctx caddy.Context) error {
	v.Handler.Provision(ctx)
	return nil
}

func (v *Vulcain) Cleanup() error {
	return v.Handler.Cleanup()
}

var (
	// _ caddy.Validator             = (*Vulcain)(nil)
	_ caddy.Provisioner         = (*Vulcain)(nil)
	_ caddy.CleanerUpper        = (*Vulcain)(nil)
	_ caddyhttp.MiddlewareHandler = (*Vulcain)(nil)
)
