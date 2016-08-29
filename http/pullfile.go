/*
* (C) 2001-2015 gtlService Inc.
*
* gutils source code
* version: 1.0.0
* author: bobliu0909@gmail.com
* datetime: 2015-10-14
*
 */

package http

import (
	"os"
	"path/filepath"
)

func HttpPullFile(savefile string, remoteurl string, resumable bool) error {

	fpath, err := filepath.Abs(savefile)
	if err != nil {
		return ErrFilePathInvalid
	}

	fd, err := os.OpenFile(fpath, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return ErrFileOpenException
	}

	httpclient := NewHttpClient(nil, 2)
	if err := httpclient.GetFile(fd, remoteurl, resumable); err != nil {
		fd.Close()
		os.Remove(fpath)
		return err
	}
	fd.Close()
	return nil
}
