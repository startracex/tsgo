package contentmappertest

import (
	"context"
	"fmt"
	"strings"

	"github.com/startracex/tsgo/pkg/contentmapper"
	"github.com/startracex/tsgo/pkg/core"
	"github.com/startracex/tsgo/pkg/json"
)

type verbatimHandler struct{ noNotifications }

type moduleVerbatimHandler struct{ noNotifications }

// formatVerbatimHandler maps content to a .cts file when it contains "cjs", and to a .mts file otherwise.
type formatVerbatimHandler struct{ noNotifications }

func (verbatimHandler) HandleRequest(ctx context.Context, method string, params json.Value) (any, error) {
	switch method {
	case contentmapper.MethodInitialize:
		return initializeResult("mapper"), nil
	case contentmapper.MethodTransform:
		var p contentmapper.TransformParams
		if err := json.Unmarshal(params, &p); err != nil {
			return nil, err
		}
		mappedOutput, err := identityMappedOutput(p.Content)
		if err != nil {
			return nil, err
		}
		return contentmapper.TransformResult{MappedOutput: mappedOutput}, nil
	default:
		return nil, fmt.Errorf("contentmappertest: unexpected method %q", method)
	}
}

func (moduleVerbatimHandler) HandleRequest(ctx context.Context, method string, params json.Value) (any, error) {
	switch method {
	case contentmapper.MethodInitialize:
		return initializeResult("mapper"), nil
	case contentmapper.MethodTransform:
		var p contentmapper.TransformParams
		if err := json.Unmarshal(params, &p); err != nil {
			return nil, err
		}
		mappedOutput, err := identityMappedOutput(p.Content)
		if err != nil {
			return nil, err
		}
		mappedOutput.Extension = ".mts"
		return contentmapper.TransformResult{MappedOutput: mappedOutput}, nil
	default:
		return nil, fmt.Errorf("contentmappertest: unexpected method %q", method)
	}
}

func (formatVerbatimHandler) HandleRequest(ctx context.Context, method string, params json.Value) (any, error) {
	switch method {
	case contentmapper.MethodInitialize:
		return initializeResult("mapper"), nil
	case contentmapper.MethodTransform:
		var p contentmapper.TransformParams
		if err := json.Unmarshal(params, &p); err != nil {
			return nil, err
		}
		mappedOutput, err := identityMappedOutput(p.Content)
		if err != nil {
			return nil, err
		}
		mappedOutput.Extension = core.IfElse(strings.Contains(p.Content, "cjs"), ".cts", ".mts")
		return contentmapper.TransformResult{MappedOutput: mappedOutput}, nil
	default:
		return nil, fmt.Errorf("contentmappertest: unexpected method %q", method)
	}
}
