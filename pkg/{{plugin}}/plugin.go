package {{plugin}}

import (
	"encoding/json"
	"reflect"
	"strconv"

	api "github.com/intel-innersource/frameworks.automation.dtac.agent/api/grpc/go"
	"github.com/intel-innersource/frameworks.automation.dtac.agent/pkg/endpoint"
	"github.com/intel-innersource/frameworks.automation.dtac.agent/pkg/plugins"
	"github.com/intel-innersource/frameworks.automation.dtac.agent/pkg/plugins/utility"
	_ "net/http/pprof" // enable remote debugging if desired
)

var _ plugins.Plugin = &Plugin{} // interface check like in the hello sample :contentReference[oaicite:2]{index=2}

type Plugin struct {
	plugins.PluginBase
	cfg Config
}

func New() *Plugin {
	p := &Plugin{
		PluginBase: plugins.PluginBase{
			Methods: make(map[string]endpoint.Func),
		},
		cfg: Config{
			// sensible defaults
			Message: "hello from {{plugin}}",
		},
	}
	p.SetRootPath("{{plugin}}") // namespacing like the sample uses "hello" :contentReference[oaicite:3]{index=3}
	return p
}

// Intentionally not pointer receiver to return the concrete type name (mirrors sample)
func (p Plugin) Name() string {
	return reflect.TypeOf(p).Name() // same technique used in hello to avoid PluginBase type leak :contentReference[oaicite:4]{index=4}
}

func (p *Plugin) Register(req *api.RegisterRequest, resp *api.RegisterResponse) error {
	*resp = api.RegisterResponse{Endpoints: make([]*api.PluginEndpoint, 0)}

	// If you need raw map[string]any you can, but typed is nicer:
	if req.Config != "" {
		var cfg Config
		if err := json.Unmarshal([]byte(req.Config), &cfg); err != nil {
			return err
		}
		// merge over defaults
		if cfg.Message != "" {
			p.cfg.Message = cfg.Message
		}
	}

	// Declare endpoints (mirrors your sample’s endpoint.New + WithOutput + Convert) :contentReference[oaicite:5]{index=5}
	authz := endpoint.AuthGroupAdmin.String()
	eps := []*endpoint.Endpoint{
		endpoint.NewEndpoint(
			"hello",
			endpoint.ActionRead,
			"returns a hello-style message",
			p.handleHello,
			req.DefaultSecure,
			authz,
			endpoint.WithOutput(&HelloOut{}),
		),
	}

	p.RegisterMethods(eps) // registers into PluginBase.Methods (same pattern) :contentReference[oaicite:6]{index=6}
	for _, ep := range eps {
		resp.Endpoints = append(resp.Endpoints, utility.ConvertEndpointToPluginEndpoint(ep)) // same conversion step :contentReference[oaicite:7]{index=7}
	}

	p.Log(plugins.LevelInfo, "{{plugin}} registered", map[string]string{
		"endpoint_count": strconv.Itoa(len(eps)),
	})

	return nil
}

type HelloOut struct {
	Message string `json:"message"`
}

// Handler mirrors the hello example’s utility wrapper approach
func (p *Plugin) handleHello(in *endpoint.Request) (*endpoint.Response, error) {
	return utility.PluginHandleWrapperWithHeaders(
		in,
		func() (map[string][]string, []byte, error) {
			headers := map[string][]string{
				"X-PLUGIN-NAME": {p.Name()},
			}
			body, err := json.Marshal(HelloOut{Message: p.cfg.Message})
			if err != nil {
				return nil, nil, err
			}
			return headers, body, nil
		},
		"{{plugin}} hello output",
	) // same utility wrapper signature used by the hello sample :contentReference[oaicite:8]{index=8}
}
