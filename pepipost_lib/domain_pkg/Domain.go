/*
 * pepipost_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package domain_pkg

import "pepipost_lib/configuration_pkg"
import "pepipost_lib/models_pkg"

/*
 * Interface for the DOMAIN_IMPL
 */
type DOMAIN interface {
    Adddomain (*models_pkg.DomainStruct) (interface{}, error)
}

/*
 * Factory for the DOMAIN interaface returning DOMAIN_IMPL
 */
func NewDOMAIN(config configuration_pkg.CONFIGURATION) *DOMAIN_IMPL {
    client := new(DOMAIN_IMPL)
    client.config = config
    return client
}