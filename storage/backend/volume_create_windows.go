//go:build windows
// +build windows

package backend

import (
	"github.com/golang/glog"
	"golang.org/x/sys/windows"

	"github.com/bookstairs/bookworm/storage/backend/memory_map"
	"github.com/bookstairs/bookworm/storage/backend/memory_map/os_overloads"
)

func CreateVolumeFile(fileName string, preallocate int64, memoryMapSizeMB uint32) (BackendStorageFile, error) {
	if preallocate > 0 {
		glog.V(0).Infof("Preallocated disk space for %s is not supported", fileName)
	}

	if memoryMapSizeMB > 0 {
		file, e := os_overloads.OpenFile(fileName, windows.O_RDWR|windows.O_CREAT, 0644, true)
		if e != nil {
			return nil, e
		}
		return memory_map.NewMemoryMappedFile(file, memoryMapSizeMB), nil
	} else {
		file, e := os_overloads.OpenFile(fileName, windows.O_RDWR|windows.O_CREAT|windows.O_TRUNC, 0644, false)
		if e != nil {
			return nil, e
		}
		return NewDiskFile(file), nil
	}

}
