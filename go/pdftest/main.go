package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// ticket 凭证信息
type ticket struct {
	name      string
	gender    string
	academy   string
	major     string
	class     string
	issueDate string
	issueTime string
	operator  string
	code      int
}

func printPDF(filename string, ti *ticket, photo io.Reader) error {
	title, subTitle := "退款凭证", "凭证信息"
	leader, department := "直接上级意见", "部门意见"
	pdf := gofpdf.New("P", "mm", "A4", "")
	pageWidth, pageHeigth := pdf.GetPageSize()
	// 设置页边距
	l, t, r := 17.8, 19.1, 17.8
	pdf.SetMargins(l, t, r)
	pdf.AddUTF8Font("sourcehanserif", "", "fonts/SourceHanSerifCN-Regular.ttf")
	pdf.AddUTF8Font("sourcehanserif", "B", "fonts/SourceHanSerifCN-Bold.ttf")
	pdf.AddPage()

	// 当前正文宽度
	contentWidth := pageWidth - l - r
	// 当前内容高度
	contentHeigth := pageHeigth - 2*t
	// 行高
	lineHeight := 12.0

	// 标题,加粗,大小16号
	pdf.SetFont("sourcehanserif", "B", 16)
	w := (pageWidth - contentWidth) / 2
	pdf.SetX(w)
	pdf.CellFormat(contentWidth, lineHeight, title, "", 0, "CM", false, 0, "")
	contentHeigth -= lineHeight
	pdf.Ln(-1)

	// 绘制边框
	pdf.SetLineWidth(0.3)
	x, y := pdf.GetXY()
	pdf.Rect(x, y, contentWidth, contentHeigth, "")
	// 其他线条宽度
	pdf.SetLineWidth(0.1)

	// 表格头,加粗,大小12号
	pdf.SetFont("sourcehanserif", "", 12)
	pdf.CellFormat(contentWidth, lineHeight, subTitle, "B", 0, "CM", false, 0, "")
	pdf.Ln(-1)
	contentHeigth -= lineHeight

	x, y = pdf.GetXY()
	photoWidth := contentWidth * 3 / 10
	photoHeight := lineHeight * 5
	pdf.SetLineWidth(0.1)
	pdf.SetFont("sourcehanserif", "", 11)
	pdf.CellFormat(photoWidth, photoHeight, "照片", "R", 0, "CM", false, 0, "")
	if photo != nil {
		// 添加图片
		photoname := "photo"
		pothoInfo := pdf.RegisterImageOptionsReader(photoname, gofpdf.ImageOptions{ImageType: "jpg"}, photo)
		if pothoInfo != nil {
			pdf.ImageOptions(photoname, x+0.5, y+0.5, photoWidth-1.5, photoHeight-1.5, false,
				gofpdf.ImageOptions{ReadDpi: true}, 0, "")
		}
	}
	tableX := pdf.GetX()
	contentHeigth -= photoHeight

	infoWidth := (contentWidth - photoWidth) / 2
	labelWidth := 15.0
	// labelHeight := photoHeight / 5
	labelHeight := lineHeight

	pdf.SetFont("sourcehanserif", "", 10)
	pdf.SetLineWidth(0.1)
	pdf.CellFormat(labelWidth, labelHeight, "姓名", "R", 0, "CM", false, 0, "")
	pdf.CellFormat(infoWidth-labelWidth, labelHeight, ti.name, "R", 0, "CM", false, 0, "")
	pdf.CellFormat(labelWidth*2, labelHeight, "性别", "R", 0, "CM", false, 0, "")
	pdf.CellFormat(infoWidth-labelWidth*2, labelHeight, ti.gender, "", 0, "CM", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(tableX)

	pdf.CellFormat(labelWidth, labelHeight, "院系", "TR", 0, "CM", false, 0, "")
	pdf.CellFormat(infoWidth-labelWidth, labelHeight, ti.academy, "TR", 0, "CM", false, 0, "")
	pdf.CellFormat(labelWidth*2, labelHeight, "领取日期", "TR", 0, "CM", false, 0, "")
	pdf.CellFormat(infoWidth-labelWidth*2, labelHeight, ti.issueDate, "T", 0, "CM", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(tableX)

	pdf.CellFormat(labelWidth, labelHeight, "专业", "TR", 0, "CM", false, 0, "")
	pdf.CellFormat(infoWidth-labelWidth, labelHeight, ti.major, "TR", 0, "CM", false, 0, "")
	pdf.CellFormat(labelWidth*2, labelHeight, "领取时间", "TR", 0, "CM", false, 0, "")
	pdf.CellFormat(infoWidth-labelWidth*2, labelHeight, ti.issueTime, "T", 0, "CM", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(tableX)

	pdf.CellFormat(labelWidth, labelHeight, "班级", "TR", 0, "CM", false, 0, "")
	pdf.CellFormat(infoWidth-labelWidth, labelHeight, ti.class, "TR", 0, "CM", false, 0, "")
	pdf.CellFormat(labelWidth*2, labelHeight, "发放人", "TR", 0, "CM", false, 0, "")
	pdf.CellFormat(infoWidth-labelWidth*2, labelHeight, ti.operator, "T", 0, "CM", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(tableX)

	pdf.CellFormat(labelWidth, labelHeight, "No", "TR", 0, "CM", false, 0, "")
	pdf.CellFormat(contentWidth-photoWidth-labelWidth, labelHeight, fmt.Sprintf("2019%05d", ti.code), "T", 0, "CM", false, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("sourcehanserif", "B", 12)
	pdf.SetLineWidth(0.1)
	pdf.CellFormat(contentWidth, lineHeight, leader, "TB", 0, "LM", false, 0, "")
	pdf.Ln(-1)
	contentHeigth -= lineHeight * 2

	y = pdf.GetY()
	pdf.SetFont("sourcehanserif", "", 11)
	labelHeight = 12
	pdf.SetXY(-80, y+contentHeigth/2-3*labelHeight)
	pdf.CellFormat(45, labelHeight, "签字:", "", 0, "L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetXY(-80, y+contentHeigth/2-2*labelHeight)
	pdf.CellFormat(45, labelHeight, "盖章:", "", 0, "L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetXY(-80, y+contentHeigth/2-1.1*labelHeight)
	pdf.CellFormat(45, labelHeight,
		fmt.Sprintf("日期:%s年%s月%s日",
			strings.Repeat(" ", 10),
			strings.Repeat(" ", 8), strings.Repeat(" ", 8)),
		"", 0, "L", false, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("sourcehanserif", "B", 12)
	pdf.SetLineWidth(0.1)
	pdf.CellFormat(contentWidth, lineHeight, department, "TB", 0, "LM", false, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("sourcehanserif", "", 11)
	pdf.SetXY(-80, -t-3*labelHeight)
	pdf.CellFormat(45, labelHeight, "签字:", "", 0, "L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetXY(-80, -t-2*labelHeight)
	pdf.CellFormat(45, labelHeight, "盖章:", "", 0, "L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetXY(-80, -t-1.1*labelHeight)
	pdf.CellFormat(45, labelHeight,
		fmt.Sprintf("日期:%s年%s月%s日",
			strings.Repeat(" ", 10),
			strings.Repeat(" ", 8), strings.Repeat(" ", 8)),
		"", 0, "L", false, 0, "")

	return pdf.OutputFileAndClose(filename)
}

func main() {
	var (
		photoPath = `0190721025425.jpg`
		filename  = "pdftest.pdf"

		ti = &ticket{
			name:      "张三",
			gender:    "男",
			academy:   "计算机学院",
			major:     "计算机科学与技术",
			class:     "07级1班",
			issueDate: "2019-09-01",
			issueTime: "10:01:00",
			operator:  "李四",
			code:      1,
		}
	)
	f, err := os.Open(photoPath)
	if err != nil {
		log.Printf("open the file of photo failed:%s", err)
	}
	if err = printPDF(filename, ti, f); err != nil {
		log.Fatalln(err)
	}
	log.Println("done")
}
