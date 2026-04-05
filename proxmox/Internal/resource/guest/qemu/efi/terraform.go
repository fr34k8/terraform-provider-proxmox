package efi

import (
	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Terraform(config *pveSDK.EfiDisk, d *schema.ResourceData) {
	if config == nil {
		d.Set(Root, nil)
		return
	}
	setting := map[string]any{
		schemaFormat:          config.Format.String(),
		schemaPreEnrolledKeys: *config.PreEnrolledKeys,
		schemaStorage:         config.Storage.String(),
	}
	if config.Type != nil {
		setting[schemaEfiType] = config.Type.String()
	}
	settings := make([]any, 1)
	settings[0] = setting
	d.Set(Root, settings)
}
