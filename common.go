/**
 * =========================================================================== *
 * define all interface in this file
 * =========================================================================== *
 */
package eventcar

import (
	"github.com/1-bi/log-api"
	"github.com/bwmarrin/snowflake"
	"os"
	"strconv"
	"strings"
	"time"
)

/**
 * --------------------------------- *
 * constant region
 * --------------------------------- *
 */
const (
	PREFIX = "eventcar"

	STATUS_DONE    int8 = 5
	STATUS_RUNNING int8 = 1
	STATUS_CANNEL  int8 = 3

	REQ_TIMEOUT = 3 * time.Second

	// ---- value is for  interface "FutureReturnResult"
	NONE         int8 = 0
	ALL_COMPLETE int8 = 1
	ANY_ERRORS   int8 = 2
	ALL_ERRORS   int8 = 3
)

//  --- check the node id propertes
func createNodeIdFile(path string, nodeNum int64) string {

	// --- get the path ---
	lastFolder := strings.LastIndex(path, "/")
	parentFolder := path[0:lastFolder]

	existed, err := pathExists(parentFolder)

	if err != nil {
		logapi.GetLogger("serviebus.common").Debug(err.Error(), nil)
	}

	if !existed {
		os.MkdirAll(parentFolder, os.ModePerm)
	}

	// --- generate file and id
	f, err := os.Create(path)

	defer f.Close()

	if err != nil {
		logapi.GetLogger("serviebus.common").Debug(err.Error(), nil)
	}

	node, err := snowflake.NewNode(nodeNum)

	if err != nil {
		logapi.GetLogger("serviebus.common").Debug(err.Error(), nil)
	}

	// Generate a snowflake ID.
	id := node.Generate()
	locIndex, err := f.WriteString(id.String())

	f.Sync()

	logapi.GetLogger("serviebus.common").Debug("Location index is "+strconv.Itoa(locIndex), nil)

	var nodeId = id.String()
	return nodeId
}

// pathExists file exist
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
