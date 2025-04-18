package total_tls_test

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testTotalTLS_Enable(rnd, zoneID string) string {
	return acctest.LoadTestCase("totaltls.tf", rnd, zoneID)
}

func testTotalTLS_Disable(rnd, zoneID string) string {
	return acctest.LoadTestCase("totaltls_disable.tf", rnd, zoneID)
}

func TestAccCloudflareTotalTLS_Enable(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_total_tls." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testTotalTLS_Enable(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "true"),
					resource.TestCheckResourceAttr(name, "certificate_authority", "google"),
				),
			},
			// {
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// 	ResourceName:      name,
			// },
		},
	})
}

func TestAccCloudflareTotalTLS_Disable(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_total_tls." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testTotalTLS_Disable(rnd, zoneID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "false"),
					resource.TestCheckResourceAttr(name, "certificate_authority", "google"),
				),
			},
			// {
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// 	ResourceName:      name,
			// },
		},
	})
}
