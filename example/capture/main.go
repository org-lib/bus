package main

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image"
	"image/png"
	"os"
)

func savePng(filename string, img image.Image) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(fmt.Sprintf("创建文件失败, filename:%s err:%s", filename, err))
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(fmt.Sprintf("关闭文件失败, filename:%s err:%s", filename, err))
		}
	}(file)
	err = png.Encode(file, img)
	if err != nil {
		fmt.Println(fmt.Sprintf("编码PNG失败, filename:%s err:%s", filename, err))
		return err
	}
	return nil
}

func screenshotAllDisplay() {
	//截取所有显示器
	n := screenshot.NumActiveDisplays() //获取显示器的数量
	fmt.Println(fmt.Sprintf("共有%d个显示器", n))
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i) //获取显示器的分辨率
		fmt.Println(fmt.Sprintf("第%d个显示器的分辨率是:%v", i+1, bounds))

		img, err := screenshot.CaptureRect(bounds) //截取整个屏幕
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		err = savePng(fileName, img)
		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("第%d个显示器截图, 保存到:%s", i+1, fileName))
	}

}
func capture(x, y, width, height int) {
	img, err := screenshot.Capture(x, y, width, height)
	if err != nil {
		return
	}
	fileName := fmt.Sprintf("test.png")
	err = savePng(fileName, img)
	if err != nil {
		panic(err)
	}
}

func main() {
	screenshotAllDisplay()
	capture(10, 10, 500, 500)

}
