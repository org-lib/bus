package main

import "github.com/org-lib/bus/draw"

func main() {
	var xyx *draw.Poster
	xyx = &draw.Poster{
		File:   "out1.png",
		Width:  1000,
		Height: 400,

		Title: "昨日慢日志汇总统计",
		X0:    80,
		Y0:    33,
		Size0: 6,

		SubTitle: "--DBA",
		X1:       320,
		Y1:       80,
		Size1:    4,

		SubCurrentDate: "当前时间",
		X2:             80,
		Y2:             100,
		Size2:          4,

		SubDateDesc: "分析时间段",
		X3:          80,
		Y3:          130,
		Size3:       4,

		SubDate1: "开始时间",
		X4:       80,
		Y4:       170,
		Size4:    4,

		SubDate2: "结束时间",
		X5:       80,
		Y5:       200,
		Size5:    4,

		Rows:  1000,
		Size6: 3,
		Y6:    30,
	}
	draw.Png(xyx)
}
