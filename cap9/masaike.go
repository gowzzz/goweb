/*
( 1 ）通过扫描图片目录，并使用图片的文件名作为键、图片的平均颜色作为值，构建出一个
由资砖图片组成的散列，也就是一个资砖图片数据库 。 通过计算图片中每个像素红 、 绿、蓝 3 种
颜色的总和，并将它们除以像素的总数量，我们就得到了一个三元组，而这个三元组就是图片的
平均颜色。.
( 2 ）根据瓷砖图片的大小，将目标图片切割成一系列尺寸更小的子图片 。

( 3 ）对于目标图片切割出的每张子图片，将它们位于左上方的第一个像素设定为该图片的平均颜色。
( 4 ）根据子图片的平均颜色，在瓷砖图片数据库中找出一张平均颜色与之最为接近的程砖图
片，然后在目标图片的相应位置上使用瓷砖图片去代替原有的子图片 。 为了找出最接近的平均颜
色，程序需要将子图片的平均颜色以及资砖图片的平均颜色都转换成三维空间中的一个点， 并计
算这两点之间的欧几里得距离 。
( 5 ）当一张瓷砖图片被选中之后，程序就会把这张图片从资砖图片数据库中移除，以此来保
证马赛克图片中的每张资砖图片都是独一无二、各不相同的 。
*/
package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

/*
averageColor 函数会把给定图片的每个像素 中的红 、绿、蓝 3 种颜色相加起来，并将这
些颜色的总和除以图片的像素数量，最后把除法计算的结果记录在一个新创建的三元组里面（这
个三元组使用包含 3 个元素的数组表示）。
*/
func averageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()
	r, g, b := 0.0, 0.0, 0.0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}
	totalPixels := float64(bounds.Max.X * bounds.Max.Y)
	return [3]float64{(r / totalPixels), (g / totalPixels), (b / totalPixels)}

}

/*
把图片缩放至指定的宽度
*/
func resizes(in image.Image, newWidth int) image.NRGBA {
	bounds := in.Bounds()
	ratio := bounds.Dx() / newWidth
	out := image.NewNRGBA(image.Rect(
		bounds.Min.X/ratio,
		bounds.Min.Y/ratio,
		bounds.Max.X/ratio,
		bounds.Max.Y/ratio,
	))

	for y, j := bounds.Min.Y, bounds.Min.Y; y < bounds.Max.Y; y, j = y+ratio, j+1 {
		for x, i := bounds.Min.X, bounds.Min.X; x < bounds.Max.X; x, i = x+ratio, x+1 {
			r, g, b, a := in.At(x, y).RGBA()
			out.SetNRGBA(i, j, color.NRGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
		}
	}

	return *out

}

/*
扫描资砖图片所在的目录来创建一个瓷砖图片数据库 。
*/
func tilesDB() map[string][3]float64 {
	fmt.Println("start populating tiles db...")

	db := make(map[string][3]float64)
	files, _ := ioutil.ReadDir("tiles")
	for _, f := range files {
		name := "tiles/" + f.Name()
		file, err := os.Open(name)
		if err == nil {
			img, _, err := image.Decode(file)
			if err == nil {
				db[name] = averageColor(img)
			} else {
				fmt.Println("err in populating:", err, name)
			}
		} else {
			fmt.Println("acnnot open file:", err, name)
		}
		file.Close()
	}
	fmt.Println("finished poplating tiles db.")
	return db
}

func sq(n float64) float64 {
	return n * n
}

/*欧几里得距离*/
func distance(p1 [3]float64, p2 [3]float64) float64 {
	return math.Sqrt(sq(p2[0]-p1[0]) + sq(p2[1]-p1[1]) + sq(p2[2]-p1[2]))
}

/*寻找*/
func nearest(target [3]float64, db *map[string][3]float64) string {
	var filename string
	smallest := 1000000.0
	for k, v := range *db {
		dist := distance(target, v)
		if dist < smallest {
			filename, smallest = k, dist
		}
	}
	delete(*db, filename)
	return filename
}

var TILESDB map[string][3]float64

func cloneTilesDB() map[string][3]float64 {
	db := make(map[string][3]float64)
	for k, v := range TILESDB {
		db[k] = v
	}
	return db
}

func upload(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("upload.html")
	t.Execute(w, nil)
}

func mosaic(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()
	r.ParseMultipartForm(10485760)
	file, _, _ := r.FormFile("image")
	defer file.Close()
	tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))
	original, _, _ := image.Decode(file)
	bounds := original.Bounds()

	newImage := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))

	db := cloneTilesDB()

	sp := image.Point{0, 0}

	for y := bounds.Min.Y; y < bounds.Max.Y; y = y + tileSize {
		for x := bounds.Min.X; x < bounds.Max.X; x = x + tileSize {
			r, g, b, _ := original.At(x, y).RGBA()
			color := [3]float64{float64(r), float64(g), float64(b)}
			nearest := nearest(color, &db)
			file, err := os.Open(nearest)
			if err == nil {
				img, _, err := image.Decode(file)
				if err == nil {
					t := resizes(img, tileSize)
					tile := t.SubImage(t.Bounds())
					tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)
					draw.Draw(newImage, tileBounds, tile, sp, draw.Src)
				} else {
					fmt.Println("err:", err, nearest)
				}
			} else {
				fmt.Println("err:", nearest)
			}
			file.Close()
		}
	}

	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	buf2 := new(bytes.Buffer)
	jpeg.Encode(buf2, newImage, nil)
	mosaic := base64.StdEncoding.EncodeToString(buf2.Bytes())

	t1 := time.Now()
	images := map[string]string{
		"original": originalStr,
		"mosaic":   mosaic,
		"duration": fmt.Sprintf("%v", t1.Sub(t0)),
	}

	t, _ := template.ParseFiles("results.html")
	t.Execute(w, images)

}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	TILESDB = tilesDB()
	fmt.Println("mosaic server started")
	server.ListenAndServe()
}
