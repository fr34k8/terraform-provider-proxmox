package powerstate

import (
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	LegacyRoot = "vm_state"

	legacyEnumStarted = "started"
)

func SchemaLegacy() *schema.Schema {
	return &schema.Schema{
		Type:          schema.TypeString,
		Deprecated:    "Use " + Root + " instead",
		Description:   "The state of the VM (" + enumRunning + ", " + legacyEnumStarted + ", " + enumStopped + ")",
		Optional:      true,
		ConflictsWith: []string{Root},
		ValidateDiagFunc: func(i any, path cty.Path) diag.Diagnostics {
			if v, ok := i.(string); ok {
				switch v {
				case enumRunning, enumStopped, legacyEnumStarted:
					return nil
				}
			}
			return diag.Diagnostics{
				diag.Diagnostic{
					Detail:   LegacyRoot + " must be one of '" + enumRunning + "', '" + enumStopped + "', '" + legacyEnumStarted + "'",
					Summary:  "invalid " + LegacyRoot,
					Severity: diag.Error},
			}
		},
		DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			return new == legacyEnumStarted
		}}
}
