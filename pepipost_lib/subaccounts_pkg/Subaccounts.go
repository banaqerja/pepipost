/*
 * pepipost_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package subaccounts_pkg

import "pepipost_lib/configuration_pkg"
import "pepipost_lib/models_pkg"

/*
 * Interface for the SUBACCOUNTS_IMPL
 */
type SUBACCOUNTS interface {
    UpdateSubaccountsPATCH (*models_pkg.Enableordisablesubacoount) (interface{}, error)
}

/*
 * Factory for the SUBACCOUNTS interaface returning SUBACCOUNTS_IMPL
 */
func NewSUBACCOUNTS(config configuration_pkg.CONFIGURATION) *SUBACCOUNTS_IMPL {
    client := new(SUBACCOUNTS_IMPL)
    client.config = config
    return client
}