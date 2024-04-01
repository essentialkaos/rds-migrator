package app

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"time"

	"github.com/essentialkaos/ek/v12/timeutil"
)

// backup backups original meta to tar.gz archive
func backup(backupName, dir string, files []string) error {
	fd, err := os.OpenFile(dir+"/"+backupName, os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		return err
	}

	defer fd.Close()

	gw := gzip.NewWriter(fd)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	for _, file := range files {
		err := backupFile(tw, file)

		if err != nil {
			return err
		}
	}

	return nil
}

// backupFile writes file data to archive
func backupFile(tw *tar.Writer, file string) error {
	fd, err := os.OpenFile(file, os.O_RDONLY, 0)

	if err != nil {
		return err
	}

	defer fd.Close()

	stat, err := fd.Stat()

	if err != nil {
		return nil
	}

	header, err := tar.FileInfoHeader(stat, "")

	if err != nil {
		return err
	}

	err = tw.WriteHeader(header)

	if err != nil {
		return err
	}

	_, err = io.Copy(tw, fd)

	return err
}

// getBackupName generates name for backup file
func getBackupName() string {
	return timeutil.Format(time.Now(), "backup_%Y%m%d_%H%M%S.tar.gz")
}
