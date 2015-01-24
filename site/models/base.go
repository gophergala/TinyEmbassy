/*
* @Author: souravray
* @Date:   2014-10-21 03:35:32
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-24 10:41:24
 */

package models

import (
	"errors"
	"regexp"
)

var (
	EmailRegexp = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)

var (
	ErrInvalidEmail = errors.New(`Invalid email`)
	ErrNotFilled    = errors.New(`Blank email`)
)

func init() {

}
