package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/weblair/ag7if/config"
	_ "github.com/weblair/ag7if/database"
	_ "github.com/weblair/ag7if/docs"
)

// BaseVersion is the SemVer-formatted string that defines the current version of ag7if.
// Build information will be added at compile-time.
const BaseVersion = "0.1.0-develop"

// BuildTime is a timestamp of when the build is run. This variable is set at compile-time.
var BuildTime string

// GitRevision is the current Git commit ID. If the tree is dirty at compile-time, an "x-" is prepended to the hash.
// This variable is set at compile-time.
var GitRevision string

// GitBranch is the name of the active Git branch at compile-time. This variable is set at compile-time.
var GitBranch string

// @title weblair ag7if
// @version 0.1.0+0
// @description UPDATE DESCRIPTION FIELD

// @contact.name UPDATE CONTACT NAME
// @contact.email UPDATE CONTACT EMAIL

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

// @host UPDATE HOST
// @BasePath /api/v1
func main() {
	version := fmt.Sprintf(
		"%s+%s.%s.%s",
		BaseVersion,
		GitBranch,
		GitRevision,
		BuildTime,
	)

	router := newRouter(version, logrus.New())
	_ = router.Run(viper.GetString("url"))
}
