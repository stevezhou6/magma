/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkWifiConfigs Wifi configuration for a network
// swagger:model network_wifi_configs
type NetworkWifiConfigs struct {

	// additional props
	AdditionalProps map[string]string `json:"additional_props,omitempty"`

	// mgmt vpn enabled
	MgmtVpnEnabled bool `json:"mgmt_vpn_enabled,omitempty"`

	// mgmt vpn proto
	MgmtVpnProto string `json:"mgmt_vpn_proto,omitempty"`

	// mgmt vpn remote
	MgmtVpnRemote string `json:"mgmt_vpn_remote,omitempty"`

	// openr enabled
	OpenrEnabled bool `json:"openr_enabled,omitempty"`

	// ping host list
	// Unique: true
	PingHostList []string `json:"ping_host_list"`

	// ping num packets
	PingNumPackets int32 `json:"ping_num_packets,omitempty"`

	// ping timeout secs
	PingTimeoutSecs int32 `json:"ping_timeout_secs,omitempty"`

	// vl auth server addr
	VlAuthServerAddr string `json:"vl_auth_server_addr,omitempty"`

	// vl auth server port
	VlAuthServerPort int32 `json:"vl_auth_server_port,omitempty"`

	// vl auth server shared secret
	VlAuthServerSharedSecret string `json:"vl_auth_server_shared_secret,omitempty"`

	// xwf config
	XwfConfig string `json:"xwf_config,omitempty"`

	// xwf dhcp dns1
	XwfDhcpDns1 string `json:"xwf_dhcp_dns1,omitempty"`

	// xwf dhcp dns2
	XwfDhcpDns2 string `json:"xwf_dhcp_dns2,omitempty"`

	// xwf partner name
	XwfPartnerName string `json:"xwf_partner_name,omitempty"`

	// xwf radius acct port
	XwfRadiusAcctPort int32 `json:"xwf_radius_acct_port,omitempty"`

	// xwf radius auth port
	XwfRadiusAuthPort int32 `json:"xwf_radius_auth_port,omitempty"`

	// xwf radius server
	XwfRadiusServer string `json:"xwf_radius_server,omitempty"`

	// xwf radius shared secret
	XwfRadiusSharedSecret string `json:"xwf_radius_shared_secret,omitempty"`

	// xwf uam secret
	XwfUamSecret string `json:"xwf_uam_secret,omitempty"`
}

// Validate validates this network wifi configs
func (m *NetworkWifiConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePingHostList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkWifiConfigs) validatePingHostList(formats strfmt.Registry) error {

	if swag.IsZero(m.PingHostList) { // not required
		return nil
	}

	if err := validate.UniqueItems("ping_host_list", "body", m.PingHostList); err != nil {
		return err
	}

	for i := 0; i < len(m.PingHostList); i++ {

		if err := validate.MinLength("ping_host_list"+"."+strconv.Itoa(i), "body", string(m.PingHostList[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkWifiConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkWifiConfigs) UnmarshalBinary(b []byte) error {
	var res NetworkWifiConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
