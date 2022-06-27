package draw

import (
	"fmt"
	"github.com/golang/freetype"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

type Poster struct {
	// image file name
	File string `json:"-"`
	// imag size
	Width  int `json:"-"`
	Height int `json:"-"`
	// text title
	Title string `json:"-"`
	X0    int
	Y0    int
	Size0 float64
	// sub title
	SubTitle string `json:"-"`
	X1       int
	Y1       int
	Size1    float64
	//
	SubCurrentDate string `json:"-"`
	X2             int
	Y2             int
	Size2          float64
	//
	SubDateDesc string `json:"-"`
	X3          int
	Y3          int
	Size3       float64
	// title of date
	SubDate1 string `json:"-"`
	X4       int
	Y4       int
	Size4    float64
	SubDate2 string `json:"-"`
	X5       int
	Y5       int
	Size5    float64
	// text context
	// only 46 max on rows
	Rows  int `json:"-"`
	Size6 float64
	Y6    int
}

func Png(p *Poster) {
	// 内容其实位置
	nums := 40 + p.Y5
	//图片的宽度
	srcWidth := p.Width
	//图片的高度
	srcHeight := nums + (p.Rows * p.Y6)
	imgfile, _ := os.Create(p.File)
	defer imgfile.Close()

	img := image.NewRGBA(image.Rect(0, 0, srcWidth, srcHeight))

	//为背景图片设置颜色
	for y := 0; y < srcWidth; y++ {
		for x := 0; x < srcHeight; x++ {
			// 黑色
			//img.Set(x, y, color.RGBA{R: 0, G: 0, B: 0, A: 255})
			// 红色
			//img.Set(x, y, color.RGBA{R: 255, G: 0, B: 0, A: 255})
			// 透明
			img.Set(x, y, color.RGBA{R: 0, G: 0, B: 0, A: 0})
		}
	}

	fontBytes, err := ioutil.ReadFile("/Users/yuandeqiao/Documents/gomodworkspace/bus/example/pngdraw/f1.ttf")
	if err != nil {
		log.Println(err)
	}
	//载入字体数据
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println("载入字体失败", err)
	}
	f := freetype.NewContext()
	//设置分辨率
	f.SetDPI(500)
	//设置字体
	f.SetFont(font)
	//设置尺寸
	f.SetClip(img.Bounds())
	//设置输出的图片
	f.SetDst(img)
	//设置字体颜色(红色)
	f.SetSrc(image.NewUniform(color.RGBA{R: 0, G: 0, B: 0, A: 255}))

	//设置标题位置
	f.SetFontSize(p.Size0)
	pt := freetype.Pt(p.X0, p.Y0)
	_, err = f.DrawString(p.Title, pt)
	//
	f.SetFontSize(p.Size1)
	pt = freetype.Pt(p.X1, p.Y1)
	_, err = f.DrawString(p.SubTitle, pt)
	//
	pt = freetype.Pt(p.X2, p.Y2)
	_, err = f.DrawString(p.SubCurrentDate, pt)
	//
	f.SetFontSize(p.Size3)
	pt = freetype.Pt(p.X3, p.Y3)
	_, err = f.DrawString(p.SubDateDesc, pt)
	//
	f.SetFontSize(p.Size4)
	pt = freetype.Pt(p.X4, p.Y4)
	_, err = f.DrawString(p.SubDate1, pt)
	f.SetFontSize(p.Size5)
	pt = freetype.Pt(p.X5, p.Y5)
	_, err = f.DrawString(p.SubDate1, pt)

	//// 设置内容位置
	f.SetFontSize(p.Size6)
	for i := 0; i < p.Rows; i++ {
		// 从低 i+3 行开始显示文本
		pt = freetype.Pt(p.X0, (i*p.Y6)+nums)
		_, err = f.DrawString(fmt.Sprintf(`%v企业微信推送图v企业微信推送图v企业微信推送图v企业微信推送图，`, i), pt)
		if err != nil {
			log.Fatal(err)
		}
	}
	//
	//以png 格式写入文件
	err = png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}
