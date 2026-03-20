package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type VideoInfoApiResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Ttl     int              `json:"ttl"`
	Data    VideoInfoApiData `json:"data"`
}

type VideoInfoApiData struct {
	Bvid      string                `json:"bvid"`
	Aid       int                   `json:"aid"`
	Videos    int                   `json:"videos"`
	Pic       string                `json:"pic"`
	Title     string                `json:"title"`
	Pubdate   int                   `json:"pubdate"`
	Ctime     int                   `json:"ctime"`
	Desc      string                `json:"desc"`
	Owner     VideoInfoApiOwner     `json:"owner"`
	Dimension VideoInfoApiDimension `json:"dimension"`
	Pages     []VideoInfoApiPage    `json:"pages"`
	Cid       int                   `json:"cid"`
}

type VideoInfoApiOwner struct {
	Mid  int    `json:"mid"`
	Name string `json:"name"`
	Face string `json:"face"`
}

type VideoInfoApiDimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Rotate int `json:"rotate"`
}

type VideoInfoApiPage struct {
	Cid        int                   `json:"cid"`
	Page       int                   `json:"page"`
	From       string                `json:"from"`
	Part       string                `json:"part"`
	Duration   int                   `json:"duration"`
	Dimension  VideoInfoApiDimension `json:"dimension"`
	FirstFrame string                `json:"first_frame"`
	Ctime      int                   `json:"ctime"`
}

type VideoUrlApiResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Ttl     int             `json:"ttl"`
	Data    VideoUrlApiData `json:"data"`
}

type VideoUrlApiData struct {
	Quantity          int      `json:"quantity"`
	Format            string   `json:"format"`
	Timelength        int      `json:"timelength"`
	AcceptFormat      string   `json:"accept_format"`
	AcceptDescription []string `json:"accept_description"`
	AcceptQuality     []int    `json:"accept_quality"`
	Durl              []Durl   `json:"durl"`
}

type Durl struct {
	Length int64  `json:"length"`
	Size   int64  `json:"size"`
	Url    string `json:"url"`
}

type VideoPageslistApiPage = VideoInfoApiPage

type VideoPageslistApiResponse struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Ttl     int                     `json:"ttl"`
	Data    []VideoPageslistApiPage `json:"data"`
}

func main() {
	c, _ := os.ReadFile("json/1.json")
	v := VideoInfoApiResponse{}
	err := json.Unmarshal(c, &v)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", v)
	fmt.Println("==========================")
	c, _ = os.ReadFile("json/2.json")
	err = json.Unmarshal(c, &v)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", v)
	fmt.Println("==========================")
	c, _ = os.ReadFile("json/3.json")
	err = json.Unmarshal(c, &v)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", v)
	fmt.Println("==========================")
	vv := VideoPageslistApiResponse{}
	c, _ = os.ReadFile("json/4.json")
	err = json.Unmarshal(c, &vv)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", vv)
	fmt.Println("==========================")
	vvv := VideoUrlApiResponse{}
	c, _ = os.ReadFile("json/5.json")
	err = json.Unmarshal(c, &vvv)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", vvv)
}
