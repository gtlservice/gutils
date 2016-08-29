/*
* (C) 2001-2015 gtlService Inc.
*
* gutils source code
* version: 1.0.0
* author: bobliu0909@gmail.com
* datetime: 2015-10-14
*
 */

package container

import (
	"errors"
)

var DEFAULT_STACK_MAXCOUNT = 1024
var DEFAULT_QUEUE_MAXCOUNT = 1024

var (
	ERR_STACK_FULL  = errors.New("stack is full.")
	ERR_STACK_EMPTY = errors.New("stack is empty.")
	ERR_QUEUE_FULL  = errors.New("queue is full.")
	ERR_QUEUE_EMPTY = errors.New("queue is empty.")
)
