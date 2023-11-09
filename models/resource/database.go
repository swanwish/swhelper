package resource

import (
	"github.com/swanwish/go-common/db"
	"github.com/swanwish/swhelper/common"
)

var (
	resourceDB = db.GetDBConnection(common.DB_CONFIG_ID_RESOURCE)
)
