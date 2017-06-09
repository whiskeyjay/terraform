package azurerm

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceArmAutoscaleSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmAutoscaleSettingsCreateOrUpdate,
		Read:   resourceArmAutoscaleSettingsRead,
		Update: resourceArmAutoscaleSettingsCreateOrUpdate,
		Delete: resourceArmAutoscaleSettingsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceArmAutoscaleSettingsCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceArmAutoscaleSettingsRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceArmAutoscaleSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
