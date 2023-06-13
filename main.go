package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Yüklü Sürücüler")
	fmt.Println("-----------------")

	// "/sys/class/block" dizinindeki blok aygıtlarını listele
	devices, err := ioutil.ReadDir("/sys/class/block")
	if err != nil {
		fmt.Println("Hata:", err)
		os.Exit(1)
	}

	// Her bir blok aygıtı için sürücü bilgilerini al ve yazdır
	for _, device := range devices {
		devicePath := filepath.Join("/sys/class/block", device.Name())
		driverPath := filepath.Join(devicePath, "device", "driver")

		// Sürücüyü belirlemek için "/sys/class/block/<device>/device/driver" yolunu kullan
		driver, err := filepath.EvalSymlinks(driverPath)
		if err != nil {
			fmt.Println("Hata:", err)
			continue
		}

		// Sürücü adını al
		driverName := filepath.Base(driver)

		fmt.Printf("Aygıt: %s, Sürücü: %s\n", device.Name(), driverName)
	}
}

