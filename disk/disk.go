package disk

import (
	"github.com/shirou/gopsutil/disk"
)

// DiskFree 根据目录获取空间可用大小
func DiskFree(folder string, unit string) (int64, error) {
	info, err := disk.Usage(folder)
	if err != nil {
		return 0, err
	}
	switch unit {
	case "GB":
		return int64(info.Free) / 1024 / 1024 / 1024, err
	case "MB":
		return int64(info.Free) / 1024 / 1024, err
	case "KB":
		return int64(info.Free) / 1024, err
	case "B":
		return int64(info.Free), err
	default:
		return 0, nil
	}
}
