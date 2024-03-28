package app

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"github.com/essentialkaos/ek/v12/jsonutil"
	"github.com/essentialkaos/ek/v12/passwd"
	"github.com/essentialkaos/ek/v12/rand"

	"github.com/essentialkaos/rds-migrator/meta"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// convert converts file with metadata from Redis-Split/v3 format to RDS/v1 format
func convert(file string, dry bool) error {
	rsMeta := &meta.RSV3{}
	err := jsonutil.Read(file, rsMeta)

	if err != nil {
		return err
	}

	rdsMeta, err := convertMeta(rsMeta)

	if err != nil {
		return err
	}

	if dry {
		return nil
	}

	return jsonutil.Write(file, rdsMeta, 0600)
}

// convertMeta converts metadata from Redis-Split/v3 format to RDS/v1 format
func convertMeta(m *meta.RSV3) (*meta.RDSV1, error) {
	result := &meta.RDSV1{
		MetaVersion: 1,
		Tags:        m.Tags,
		Desc:        m.Desc,
		UUID:        m.UUID,
		Compatible:  m.Compatible,
		ID:          m.ID,
		Created:     m.Created,
		Auth: &meta.RDSV1InstanceAuth{
			User:   m.AuthInfo.User,
			Pepper: m.AuthInfo.Pepper,
			Hash:   m.AuthInfo.Hash,
		},
		Config: &meta.RDSV1InstanceConfigInfo{
			Hash: m.ConfigInfo.Hash,
			Date: m.ConfigInfo.Date,
		},
		Preferencies: &meta.RDSV1InstancePreferencies{
			AdminPassword:    genPassword(),
			SyncPassword:     genPassword(),
			SentinelPassword: genPassword(),
			ServicePassword:  m.Preferencies.Password,
			IsSaveDisabled:   m.Preferencies.IsSaveDisabled,
			ReplicationType:  getReplicationType(m),
		},
	}

	return result, nil
}

// getReplicationType returns replication type
func getReplicationType(m *meta.RSV3) string {
	if m.SlaveType != "" {
		switch m.SlaveType {
		case "standby":
			return "standby"
		default:
			return "replica"
		}
	}

	return m.ReplicationType
}

// genPassword generates pseudo-secure password with random length (16-28)
func genPassword() string {
	return passwd.GenPassword(16+rand.Int(6), passwd.STRENGTH_MEDIUM)
}
