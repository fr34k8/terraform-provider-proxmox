package powerstate

import (
	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func terraformLegacy(config pveSDK.PowerState, d *schema.ResourceData) {
	d.Set(LegacyRoot, terraform(config))
}

func terraformLegacyClear(d *schema.ResourceData) {
	if _, ok := d.GetOk(LegacyRoot); ok {
		d.Set(LegacyRoot, nil)
	}
}
