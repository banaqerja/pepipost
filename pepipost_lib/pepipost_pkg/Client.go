/*
 * pepipost_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package PepipostClient

import(
	"pepipost_lib/configuration_pkg"
	"pepipost_lib/mailsend_pkg"
	"pepipost_lib/events_pkg"
	"pepipost_lib/stats_pkg"
	"pepipost_lib/subaccounts_pkg"
	"pepipost_lib/subaccountsdelete_pkg"
	"pepipost_lib/subaccountsgetsubaccounts_pkg"
	"pepipost_lib/subaccountsgetcreditddetails_pkg"
	"pepipost_lib/setrecurringcreditddetails_pkg"
	"pepipost_lib/subaccountssetsubaccountcredit_pkg"
	"pepipost_lib/subaccountsupdatesubaccount_pkg"
	"pepipost_lib/subaccountscreatesubaccount_pkg"
	"pepipost_lib/suppression_pkg"
	"pepipost_lib/domaindelete_pkg"
	"pepipost_lib/domaingetdomains_pkg"
	"pepipost_lib/template_pkg"
	"pepipost_lib/domain_pkg"
)
/*
 * Client structure as interface implementation
 */
type PEPIPOST_IMPL struct {
     mailsend mailsend_pkg.MAILSEND
     events events_pkg.EVENTS
     stats stats_pkg.STATS
     subaccounts subaccounts_pkg.SUBACCOUNTS
     subaccountsdelete subaccountsdelete_pkg.SUBACCOUNTSDELETE
     subaccountsgetsubaccounts subaccountsgetsubaccounts_pkg.SUBACCOUNTSGETSUBACCOUNTS
     subaccountsgetcreditddetails subaccountsgetcreditddetails_pkg.SUBACCOUNTSGETCREDITDDETAILS
     setrecurringcreditddetails setrecurringcreditddetails_pkg.SETRECURRINGCREDITDDETAILS
     subaccountssetsubaccountcredit subaccountssetsubaccountcredit_pkg.SUBACCOUNTSSETSUBACCOUNTCREDIT
     subaccountsupdatesubaccount subaccountsupdatesubaccount_pkg.SUBACCOUNTSUPDATESUBACCOUNT
     subaccountscreatesubaccount subaccountscreatesubaccount_pkg.SUBACCOUNTSCREATESUBACCOUNT
     suppression suppression_pkg.SUPPRESSION
     domaindelete domaindelete_pkg.DOMAINDELETE
     domaingetdomains domaingetdomains_pkg.DOMAINGETDOMAINS
     template template_pkg.TEMPLATE
     domain domain_pkg.DOMAIN
     config  configuration_pkg.CONFIGURATION
}

/**
     * Access to Configuration
     * @return Returns the Configuration instance
*/
func (me *PEPIPOST_IMPL) Configuration() configuration_pkg.CONFIGURATION {
    return me.config
}
/**
     * Access to MailSend controller
     * @return Returns the MailSend() instance
*/
func (me *PEPIPOST_IMPL) MailSend() mailsend_pkg.MAILSEND {
    if(me.mailsend) == nil {
        me.mailsend = mailsend_pkg.NewMAILSEND(me.config)
    }
    return me.mailsend
}
/**
     * Access to Events controller
     * @return Returns the Events() instance
*/
func (me *PEPIPOST_IMPL) Events() events_pkg.EVENTS {
    if(me.events) == nil {
        me.events = events_pkg.NewEVENTS(me.config)
    }
    return me.events
}
/**
     * Access to Stats controller
     * @return Returns the Stats() instance
*/
func (me *PEPIPOST_IMPL) Stats() stats_pkg.STATS {
    if(me.stats) == nil {
        me.stats = stats_pkg.NewSTATS(me.config)
    }
    return me.stats
}
/**
     * Access to Subaccounts controller
     * @return Returns the Subaccounts() instance
*/
func (me *PEPIPOST_IMPL) Subaccounts() subaccounts_pkg.SUBACCOUNTS {
    if(me.subaccounts) == nil {
        me.subaccounts = subaccounts_pkg.NewSUBACCOUNTS(me.config)
    }
    return me.subaccounts
}
/**
     * Access to SubaccountsDelete controller
     * @return Returns the SubaccountsDelete() instance
*/
func (me *PEPIPOST_IMPL) SubaccountsDelete() subaccountsdelete_pkg.SUBACCOUNTSDELETE {
    if(me.subaccountsdelete) == nil {
        me.subaccountsdelete = subaccountsdelete_pkg.NewSUBACCOUNTSDELETE(me.config)
    }
    return me.subaccountsdelete
}
/**
     * Access to SubaccountsGetSubAccounts controller
     * @return Returns the SubaccountsGetSubAccounts() instance
*/
func (me *PEPIPOST_IMPL) SubaccountsGetSubAccounts() subaccountsgetsubaccounts_pkg.SUBACCOUNTSGETSUBACCOUNTS {
    if(me.subaccountsgetsubaccounts) == nil {
        me.subaccountsgetsubaccounts = subaccountsgetsubaccounts_pkg.NewSUBACCOUNTSGETSUBACCOUNTS(me.config)
    }
    return me.subaccountsgetsubaccounts
}
/**
     * Access to SubaccountsGetcreditddetails controller
     * @return Returns the SubaccountsGetcreditddetails() instance
*/
func (me *PEPIPOST_IMPL) SubaccountsGetcreditddetails() subaccountsgetcreditddetails_pkg.SUBACCOUNTSGETCREDITDDETAILS {
    if(me.subaccountsgetcreditddetails) == nil {
        me.subaccountsgetcreditddetails = subaccountsgetcreditddetails_pkg.NewSUBACCOUNTSGETCREDITDDETAILS(me.config)
    }
    return me.subaccountsgetcreditddetails
}
/**
     * Access to Setrecurringcreditddetails controller
     * @return Returns the Setrecurringcreditddetails() instance
*/
func (me *PEPIPOST_IMPL) Setrecurringcreditddetails() setrecurringcreditddetails_pkg.SETRECURRINGCREDITDDETAILS {
    if(me.setrecurringcreditddetails) == nil {
        me.setrecurringcreditddetails = setrecurringcreditddetails_pkg.NewSETRECURRINGCREDITDDETAILS(me.config)
    }
    return me.setrecurringcreditddetails
}
/**
     * Access to SubaccountsSetsubaccountcredit controller
     * @return Returns the SubaccountsSetsubaccountcredit() instance
*/
func (me *PEPIPOST_IMPL) SubaccountsSetsubaccountcredit() subaccountssetsubaccountcredit_pkg.SUBACCOUNTSSETSUBACCOUNTCREDIT {
    if(me.subaccountssetsubaccountcredit) == nil {
        me.subaccountssetsubaccountcredit = subaccountssetsubaccountcredit_pkg.NewSUBACCOUNTSSETSUBACCOUNTCREDIT(me.config)
    }
    return me.subaccountssetsubaccountcredit
}
/**
     * Access to SubaccountsUpdateSubaccount controller
     * @return Returns the SubaccountsUpdateSubaccount() instance
*/
func (me *PEPIPOST_IMPL) SubaccountsUpdateSubaccount() subaccountsupdatesubaccount_pkg.SUBACCOUNTSUPDATESUBACCOUNT {
    if(me.subaccountsupdatesubaccount) == nil {
        me.subaccountsupdatesubaccount = subaccountsupdatesubaccount_pkg.NewSUBACCOUNTSUPDATESUBACCOUNT(me.config)
    }
    return me.subaccountsupdatesubaccount
}
/**
     * Access to SubaccountsCreateSubaccount controller
     * @return Returns the SubaccountsCreateSubaccount() instance
*/
func (me *PEPIPOST_IMPL) SubaccountsCreateSubaccount() subaccountscreatesubaccount_pkg.SUBACCOUNTSCREATESUBACCOUNT {
    if(me.subaccountscreatesubaccount) == nil {
        me.subaccountscreatesubaccount = subaccountscreatesubaccount_pkg.NewSUBACCOUNTSCREATESUBACCOUNT(me.config)
    }
    return me.subaccountscreatesubaccount
}
/**
     * Access to Suppression controller
     * @return Returns the Suppression() instance
*/
func (me *PEPIPOST_IMPL) Suppression() suppression_pkg.SUPPRESSION {
    if(me.suppression) == nil {
        me.suppression = suppression_pkg.NewSUPPRESSION(me.config)
    }
    return me.suppression
}
/**
     * Access to DomainDelete controller
     * @return Returns the DomainDelete() instance
*/
func (me *PEPIPOST_IMPL) DomainDelete() domaindelete_pkg.DOMAINDELETE {
    if(me.domaindelete) == nil {
        me.domaindelete = domaindelete_pkg.NewDOMAINDELETE(me.config)
    }
    return me.domaindelete
}
/**
     * Access to DomainGetDomains controller
     * @return Returns the DomainGetDomains() instance
*/
func (me *PEPIPOST_IMPL) DomainGetDomains() domaingetdomains_pkg.DOMAINGETDOMAINS {
    if(me.domaingetdomains) == nil {
        me.domaingetdomains = domaingetdomains_pkg.NewDOMAINGETDOMAINS(me.config)
    }
    return me.domaingetdomains
}
/**
     * Access to Template controller
     * @return Returns the Template() instance
*/
func (me *PEPIPOST_IMPL) Template() template_pkg.TEMPLATE {
    if(me.template) == nil {
        me.template = template_pkg.NewTEMPLATE(me.config)
    }
    return me.template
}
/**
     * Access to Domain controller
     * @return Returns the Domain() instance
*/
func (me *PEPIPOST_IMPL) Domain() domain_pkg.DOMAIN {
    if(me.domain) == nil {
        me.domain = domain_pkg.NewDOMAIN(me.config)
    }
    return me.domain
}
