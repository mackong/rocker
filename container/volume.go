package container

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
)

func MountVolume(rootfs, volume string) error {
	source, target, err := volumeSplit(rootfs, volume)
	if err != nil {
		return err
	}

	//不存在则创建目录
	if _, err := os.Stat(target); os.IsNotExist(err) {
		os.Mkdir(target, 0700)
	}

	if err := syscall.Mount(source, target, "", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		log.Errorf("mount volume %s => %s error ", source, target, err)
		return err
	}
	log.Info("mount volume %s done. %v", volume, err)
	return nil
}
func volumeSplit(rootfs, volume string) (source, target string, err error) {
	if volume == "" {
		return
	}

	volumeSplit := strings.Split(volume, ":")
	if len(volumeSplit) != 2 {
		err = fmt.Errorf("错误的Volume参数 %s", volume)
		return
	}
	source = volumeSplit[0]
	target = path.Join(rootfs, volumeSplit[1])
	return
}

func UnMountVolume(id, volume string) error {
	if volume == "" {
		return nil
	}

	//TODO::判断是否创建了容器内的目录，删除

	_, target, err := volumeSplit(getMergedPath(id), volume)
	if err != nil {
		return err
	}
	if _, err := exec.Command("umount", target).CombinedOutput(); err != nil {
		log.Errorf("umount volume %s failed. %v", target, err)
		return err
	}

	log.Info("umount volume %s done. %v", volume, err)
	return nil
}