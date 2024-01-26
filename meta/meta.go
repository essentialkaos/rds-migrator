package meta

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// Redis-Split v2/v3 //////////////////////////////////////////////////////////////////////

type RSV3InstancePreferencies struct {
	ID             int    `json:"id"`               // Instance ID
	Password       string `json:"password"`         // Redis password for secure instance
	Prefix         string `json:"prefix"`           // Commands prefix
	IsSecure       bool   `json:"secure"`           // Secure instance flag
	IsSaveDisabled bool   `json:"is_save_disabled"` // Disabled saves flag
}

type RSV3InstanceAuthInfo struct {
	Pepper string `json:"pepper"`
	Hash   string `json:"hash"`
	User   string `json:"user"`
}

type RSV3InstanceConfigInfo struct {
	Hash string `json:"hash"`
	Date int64  `json:"date"`
}

type RSV3 struct {
	Tags            []string                  `json:"tags,omitempty"`   // List of tags
	Desc            string                    `json:"desc"`             // Description
	SlaveType       string                    `json:"slave_type"`       // Replication type (v2)
	ReplicationType string                    `json:"replication_type"` // Replication type (v3)
	UUID            string                    `json:"uuid"`             // UUID
	Compatible      string                    `json:"compatible"`       // Compatible redis version
	MetaVersion     int                       `json:"meta_version"`     // Meta information version
	ID              int                       `json:"id"`               // Instance ID
	Created         int64                     `json:"created"`          // Date of creation (unix timestamp)
	Preferencies    *RSV3InstancePreferencies `json:"preferencies"`     // Config data
	AuthInfo        *RSV3InstanceAuthInfo     `json:"auth"`             // Instance auth info
	ConfigInfo      *RSV3InstanceConfigInfo   `json:"config"`           // Config info (hash + creation date)
	Sentinel        bool                      `json:"sentinel"`         // Sentinel monitoring flag
}

// RDS v1 //////////////////////////////////////////////////////////////////////////////

type RDSV1 struct {
	Tags         []string                   `json:"tags,omitempty"`       // List of tags
	Desc         string                     `json:"desc"`                 // Description
	UUID         string                     `json:"uuid"`                 // UUID
	Compatible   string                     `json:"compatible,omitempty"` // Compatible redis version
	MetaVersion  int                        `json:"meta_version"`         // Meta information version
	ID           int                        `json:"id"`                   // Instance ID
	Created      int64                      `json:"created"`              // Date of creation (unix timestamp)
	Preferencies *RDSV1InstancePreferencies `json:"preferencies"`         // Config data
	Config       *RDSV1InstanceConfigInfo   `json:"config"`               // Config info (hash + creation date)
	Auth         *RDSV1InstanceAuth         `json:"auth"`                 // Instance auth info
}

type RDSV1InstanceAuth struct {
	User   string `json:"user"`
	Pepper string `json:"pepper"`
	Hash   string `json:"hash"`
}

type RDSV1InstanceConfigInfo struct {
	Hash string `json:"hash"`
	Date int64  `json:"date"`
}

type RDSV1InstancePreferencies struct {
	AdminPassword    string `json:"admin_password,omitempty"`   // Admin user password
	SyncPassword     string `json:"sync_password,omitempty"`    // Sync user password
	ServicePassword  string `json:"service_password,omitempty"` // Service user password
	SentinelPassword string `json:"sentinel_password"`          // Sentinel user password
	ReplicationType  string `json:"replication_type"`           // Replication type
	IsSaveDisabled   bool   `json:"is_save_disabled"`           // Disabled saves flag
}

// ////////////////////////////////////////////////////////////////////////////////// //
