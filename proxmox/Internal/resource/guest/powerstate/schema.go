package powerstate

import (
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	Root = "power_state"

	Default string = enumRunning

	enumRunning = "running"
	enumStopped = "stopped"
)

func Schema(s schema.Schema) *schema.Schema {
	s.Type = schema.TypeString
	s.Optional = true
	s.ValidateDiagFunc = func(i any, path cty.Path) diag.Diagnostics {
		if v, ok := i.(string); ok {
			switch v {
			case enumRunning, enumStopped:
				return nil
			}
		}
		return diag.Diagnostics{
			diag.Diagnostic{
				Detail:   "the power state must be either '" + enumRunning + "' or '" + enumStopped + "'",
				Summary:  "invalid power state",
				Severity: diag.Error},
		}
	}
	return &s
}
