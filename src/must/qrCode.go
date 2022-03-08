package must

import (
	"fmt"
	"image/color"
	"image/png"

	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/tuotoo/qrcode"

	code "github.com/skip2/go-qrcode"
)

func init() {
	//GetQrCode()
	//fi,err := os.Open("qr1.png")
	//if err !=nil{
	//	fmt.Println("读文件错了",err)
	//	return
	//}
	//defer fi.Close()
	//ReadQrCode(fi)
	//GetSkip2QrCode()
}

// 获取一张二维码
func GetQrCode() {
	// 生成二维码信息
	qrCode, _ := qr.Encode("lalalala", qr.H, qr.Auto)

	// 设定二维码尺寸
	qrCode, _ = barcode.Scale(qrCode, 256, 256)

	// 创建一个Png文件
	file, _ := os.Create("qr1.png")
	defer file.Close()

	// 将内容传入Png文件
	png.Encode(file, qrCode)
}

// 读取一张二维码
func ReadQrCode(fi *os.File) {
	// 读取一张
	qrMatrix, err := qrcode.Decode(fi)
	if err != nil {
		fmt.Println("读取二维码错误：", err.Error())
		return
	}
	fmt.Println(qrMatrix.Content)
}

// skip2获取二维码
func GetSkip2QrCode() {
	// 生成一张字节二维码
	png, err := code.Encode("hello", code.Medium, 256)
	if err != nil {
		fmt.Println("生成二维码字节失败")
	}
	// 生成二维码->到指定地址
	err = code.WriteFile("hello", code.Medium, 256, "qr.png")
	if err != nil {
		fmt.Println("生成二维码到指定地址失败")
	}
	// 生成指定颜色二维码->到指定地址
	err = code.WriteColorFile("hello", code.Medium, 256, color.Black, color.White, "qr.png")
	if err != nil {
		fmt.Println("生成指定颜色二维码到指定地址失败")
	}
	fmt.Println(png)
}
