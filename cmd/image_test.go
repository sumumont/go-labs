package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-labs/internal/logging"
	"golang.org/x/image/bmp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"testing"
	"time"
)

func TestCut(t *testing.T) {
	logging.Debug().Msg("begin")
	start := time.Now().UnixNano() / 1000000

	logging.Debug().Interface("start", start).Send()
	box := Box{
		X:      30,
		Y:      630,
		Width:  500,
		Height: 500,
	}

	src := "D:\\OneDrive\\OneDrive - 依瞳科技（深圳）有限公司\\桌面\\12321.jpg"
	dst := strings.Replace(src, ".", "_cut.", 1)
	fmt.Println("src=", src, " dst=", dst)
	fIn, _ := os.Open(src)
	defer fIn.Close()

	fOut, _ := os.Create(dst)
	defer fOut.Close()

	x0 := box.X
	y0 := box.Y
	x1 := box.X + box.Width
	y1 := box.Y + box.Height
	err := Clip(fIn, fOut, 0, 0, x0, y0, x1, y1, 100)
	if err != nil {
		panic(err)
	}
	end := time.Now().UnixNano() / 1000000

	logging.Debug().Interface("end", end).Send()
	logging.Debug().Interface("time", end-start).Send()
	logging.Debug().Msg("end")
}

func Clip(in io.Reader, out io.Writer, wi, hi, x0, y0, x1, y1, quality int) (err error) {
	err = errors.New("unknow error")
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	var origin image.Image
	var fm string
	origin, fm, err = image.Decode(in)
	if err != nil {
		log.Println(err)
		return err
	}

	if wi == 0 || hi == 0 {
		wi = origin.Bounds().Max.X
		hi = origin.Bounds().Max.Y
	}
	var canvas image.Image
	canvas = origin
	switch fm {
	case "jpg":
		//img := canvas.(*image.RGBA)
		//subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
		//return png.Encode(out, subImg)
		img := canvas.(*image.YCbCr)
		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.YCbCr)
		return jpeg.Encode(out, subImg, &jpeg.Options{quality})
	case "jpeg":
		img := canvas.(*image.YCbCr)
		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.YCbCr)
		return jpeg.Encode(out, subImg, &jpeg.Options{quality})
	case "png":
		switch canvas.(type) {
		case *image.NRGBA:
			img := canvas.(*image.NRGBA)
			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.NRGBA)
			return png.Encode(out, subImg)
		case *image.RGBA:
			img := canvas.(*image.RGBA)
			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
			return png.Encode(out, subImg)
		}
	case "gif":
		img := canvas.(*image.Paletted)
		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.Paletted)
		return gif.Encode(out, subImg, &gif.Options{})
	case "bmp":
		img := canvas.(*image.RGBA)
		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
		return bmp.Encode(out, subImg)
	default:
		return errors.New("ERROR FORMAT")
	}
	return nil
}

func TestCropImage(t *testing.T) {
	var param = `{
    "aiImageName": "LWA51XXB45965782P0013_LWA51XXB45965782P0014_Img9.jpg",
    "aiImagePath": "data-sources/image/1/1/1/92af745f-b090-4361-95a1-4a5fb1ea67f8/LWA51XXB45965782P0013_LWA51XXB45965782P0014_Img9.jpg",
    "Objects": [
        {
            "box": "{\"X\":729, \"Y\":2454, \"Angle\":0, \"Width\":166, \"Height\":153, \"DefectType\":\"chajie_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 729,
                "Y": 2454,
                "Angle": 0,
                "result": "",
                "Width": 166,
                "Height": 153,
                "DefectType": "chajie_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "chajie",
            "ocr": null,
            "result": "OK",
            "score": 0.984863,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":729, \"Y\":2434, \"Angle\":0, \"Width\":166, \"Height\":21, \"DefectType\":\"lianxi_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 729,
                "Y": 2434,
                "Angle": 0,
                "result": "",
                "Width": 166,
                "Height": 21,
                "DefectType": "lianxi_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "lianxi",
            "ocr": null,
            "result": "OK",
            "score": 0.963379,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":729, \"Y\":2606, \"Angle\":0, \"Width\":166, \"Height\":21, \"DefectType\":\"lianxi_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 729,
                "Y": 2606,
                "Angle": 0,
                "result": "",
                "Width": 166,
                "Height": 21,
                "DefectType": "lianxi_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "lianxi",
            "ocr": null,
            "result": "OK",
            "score": 0.953125,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":709, \"Y\":2454, \"Angle\":0, \"Width\":21, \"Height\":153, \"DefectType\":\"lianxi_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 709,
                "Y": 2454,
                "Angle": 0,
                "result": "",
                "Width": 21,
                "Height": 153,
                "DefectType": "lianxi_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "lianxi",
            "ocr": null,
            "result": "OK",
            "score": 0.98291,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":894, \"Y\":2454, \"Angle\":0, \"Width\":21, \"Height\":153, \"DefectType\":\"lianxi_ng\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 894,
                "Y": 2454,
                "Angle": 0,
                "result": "",
                "Width": 21,
                "Height": 153,
                "DefectType": "lianxi_ng",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "lianxi",
            "ocr": null,
            "result": "NG",
            "score": 0.552246,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":427, \"Y\":2465, \"Angle\":0, \"Width\":131, \"Height\":128, \"DefectType\":\"chajie_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 427,
                "Y": 2465,
                "Angle": 0,
                "result": "",
                "Width": 131,
                "Height": 128,
                "DefectType": "chajie_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "chajie",
            "ocr": null,
            "result": "OK",
            "score": 0.919922,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":427, \"Y\":2445, \"Angle\":0, \"Width\":131, \"Height\":21, \"DefectType\":\"lianxi_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 427,
                "Y": 2445,
                "Angle": 0,
                "result": "",
                "Width": 131,
                "Height": 21,
                "DefectType": "lianxi_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "lianxi",
            "ocr": null,
            "result": "OK",
            "score": 0.922852,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":427, \"Y\":2592, \"Angle\":0, \"Width\":131, \"Height\":21, \"DefectType\":\"lianxi_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 427,
                "Y": 2592,
                "Angle": 0,
                "result": "",
                "Width": 131,
                "Height": 21,
                "DefectType": "lianxi_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "lianxi",
            "ocr": null,
            "result": "OK",
            "score": 0.935547,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":407, \"Y\":2465, \"Angle\":0, \"Width\":21, \"Height\":128, \"DefectType\":\"lianxi_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 407,
                "Y": 2465,
                "Angle": 0,
                "result": "",
                "Width": 21,
                "Height": 128,
                "DefectType": "lianxi_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "lianxi",
            "ocr": null,
            "result": "OK",
            "score": 0.969238,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":557, \"Y\":2465, \"Angle\":0, \"Width\":21, \"Height\":129, \"DefectType\":\"lianxi_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 557,
                "Y": 2465,
                "Angle": 0,
                "result": "",
                "Width": 21,
                "Height": 129,
                "DefectType": "lianxi_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "lianxi",
            "ocr": null,
            "result": "OK",
            "score": 0.947754,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":271, \"Y\":2464, \"Angle\":0, \"Width\":135, \"Height\":129, \"DefectType\":\"chajie_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 271,
                "Y": 2464,
                "Angle": 0,
                "result": "",
                "Width": 135,
                "Height": 129,
                "DefectType": "chajie_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "chajie",
            "ocr": null,
            "result": "OK",
            "score": 0.947754,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":271, \"Y\":2444, \"Angle\":0, \"Width\":135, \"Height\":21, \"DefectType\":\"lianxi_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 271,
                "Y": 2444,
                "Angle": 0,
                "result": "",
                "Width": 135,
                "Height": 21,
                "DefectType": "lianxi_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "lianxi",
            "ocr": null,
            "result": "OK",
            "score": 0.927246,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":271, \"Y\":2592, \"Angle\":0, \"Width\":135, \"Height\":21, \"DefectType\":\"lianxi_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 271,
                "Y": 2592,
                "Angle": 0,
                "result": "",
                "Width": 135,
                "Height": 21,
                "DefectType": "lianxi_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "lianxi",
            "ocr": null,
            "result": "OK",
            "score": 0.832031,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":251, \"Y\":2464, \"Angle\":0, \"Width\":21, \"Height\":129, \"DefectType\":\"lianxi_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 251,
                "Y": 2464,
                "Angle": 0,
                "result": "",
                "Width": 21,
                "Height": 129,
                "DefectType": "lianxi_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "lianxi",
            "ocr": null,
            "result": "OK",
            "score": 0.942383,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":405, \"Y\":2464, \"Angle\":0, \"Width\":21, \"Height\":129, \"DefectType\":\"lianxi_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 405,
                "Y": 2464,
                "Angle": 0,
                "result": "",
                "Width": 21,
                "Height": 129,
                "DefectType": "lianxi_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "lianxi",
            "ocr": null,
            "result": "OK",
            "score": 0.953613,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":590, \"Y\":2465, \"Angle\":0, \"Width\":126, \"Height\":126, \"DefectType\":\"chajie_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 590,
                "Y": 2465,
                "Angle": 0,
                "result": "",
                "Width": 126,
                "Height": 126,
                "DefectType": "chajie_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "chajie",
            "ocr": null,
            "result": "OK",
            "score": 0.991699,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":590, \"Y\":2445, \"Angle\":0, \"Width\":126, \"Height\":21, \"DefectType\":\"lianxi_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 590,
                "Y": 2445,
                "Angle": 0,
                "result": "",
                "Width": 126,
                "Height": 21,
                "DefectType": "lianxi_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "lianxi",
            "ocr": null,
            "result": "OK",
            "score": 0.976074,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":590, \"Y\":2590, \"Angle\":0, \"Width\":126, \"Height\":21, \"DefectType\":\"lianxi_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 590,
                "Y": 2590,
                "Angle": 0,
                "result": "",
                "Width": 126,
                "Height": 21,
                "DefectType": "lianxi_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "lianxi",
            "ocr": null,
            "result": "OK",
            "score": 0.981445,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":570, \"Y\":2465, \"Angle\":0, \"Width\":21, \"Height\":126, \"DefectType\":\"lianxi_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 570,
                "Y": 2465,
                "Angle": 0,
                "result": "",
                "Width": 21,
                "Height": 126,
                "DefectType": "lianxi_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "lianxi",
            "ocr": null,
            "result": "OK",
            "score": 0.934082,
            "segmentation": null,
            "sub_objects": null
        },
        {
            "box": "{\"X\":715, \"Y\":2465, \"Angle\":0, \"Width\":21, \"Height\":126, \"DefectType\":\"lianxi_ok\", \"DetailLabel\":\"T9101\"}",
            "boxx": {
                "X": 715,
                "Y": 2465,
                "Angle": 0,
                "result": "",
                "Width": 21,
                "Height": 126,
                "DefectType": "lianxi_ok",
                "DetailLabel": "T9101"
            },
            "classification": null,
            "label": "lianxi",
            "ocr": null,
            "result": "OK",
            "score": 0.994141,
            "segmentation": null,
            "sub_objects": null
        }
    ]
}`
	var aiResult AiResult
	err := json.Unmarshal([]byte(param), &aiResult)
	if err != nil {
		panic(err)
	}
	x0, y0, x1, y1 := math.MaxInt64, math.MaxInt64, 0, 0
	objects := aiResult.Objects
	partName := "T9101"
	for _, obj := range objects {
		boxStr := obj.Box
		var box Box
		err = json.Unmarshal([]byte(boxStr), &box)
		if err != nil {
			panic(err)
		}
		// 找到同名器件名
		if box.DetailLabel == partName {

			//最左的X
			if x0 > box.X {
				x0 = box.X
			}

			//最上的y
			if y0 > box.Y {
				y0 = box.Y
			}

			//最右的x
			if x1 < box.X+box.Width {
				x1 = box.X + box.Width
			}

			//最下的y
			if y1 < box.Y+box.Height {
				y1 = box.Y + box.Height
			}
		}
	}
	result := fmt.Sprintf("%v %v %v %v", x0, y0, x1, y1)
	fmt.Printf(result)
}

type Point struct {
	X      float64 `json:"X"`
	Y      float64 `json:"Y"`
	Width  float64 `json:"Width"`
	Height float64 `json:"Height"`
}

type Rectangle struct {
	LT Point
	LB Point
	RB Point
	RT Point
}

func (rec *Rectangle) Build(box Point) {
	rec.LT = box
	rec.LB = Point{
		X: box.X,
		Y: box.Y + box.Height,
	}

	rec.RB = Point{
		X: box.X + box.Width,
		Y: box.Y + box.Height,
	}

	rec.RT = Point{
		X: box.X + box.Width,
		Y: box.Y,
	}
}
func TestUIO(t *testing.T) {

	boxA := Point{
		X:      100,
		Y:      100,
		Width:  100,
		Height: 100,
	}
	Ra := Rectangle{}
	Ra.Build(boxA)

	boxB := Point{
		X:      50,
		Y:      50,
		Width:  10,
		Height: 10,
	}
	Rb := Rectangle{}
	Rb.Build(boxB)
	fmt.Println(Iou(Ra, Rb))
}
func Iou(A, B Rectangle) float64 {
	W := math.Min(A.RT.X, B.RT.X) - math.Max(A.LB.X, B.LB.X)
	H := math.Min(A.RB.Y, B.RB.Y) - math.Max(A.LT.Y, B.LT.Y)
	if W <= 0 || H <= 0 {
		return 0
	}
	SA := math.Abs(A.RT.X-A.LB.X) * math.Abs(A.RT.Y-A.LB.Y)
	SB := math.Abs(B.RT.X-B.LB.X) * math.Abs(B.RT.Y-B.LB.Y)
	cross := W * H
	return cross / (SA + SB - cross)
}
