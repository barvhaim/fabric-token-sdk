/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sqlite

import (
	"database/sql"

	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/db/driver/sql/sqlite"
	"github.com/hyperledger-labs/fabric-token-sdk/token/services/db/driver"
	"github.com/hyperledger-labs/fabric-token-sdk/token/services/db/sql/common"
)

func OpenWalletDB(k common.Opts) (driver.WalletDB, error) {
	db, err := sqlite.OpenDB(k.DataSource, k.MaxOpenConns, k.SkipPragmas)
	if err != nil {
		return nil, err
	}
	return NewWalletDB(db, common.NewDBOptsFromOpts(k))
}

func NewWalletDB(db *sql.DB, opts common.NewDBOpts) (driver.WalletDB, error) {
	return common.NewWalletDB(db, opts)
}
