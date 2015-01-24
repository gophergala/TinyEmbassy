/*
* @Author: souravray
* @Date:   2015-01-24 11:30:37
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-24 11:32:02
 */

package models

import (
	"errors"
)

var (
	EmailRegexp = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)

var (
	ErrInvalidReferer = errors.New(`Invalid referer`)
)

func init() {

}
