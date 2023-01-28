package main

/**
 * @title main
 * @author CH00SE1
 * @date 2022-10-15 13:38:32
 */

import (
	. "encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"test/entiy"
)

const (
	Url       = "https://www.appmovie.cc"
	VideoUrl  = "https://www.appmovie.cc/index.php/vod/detail/id/461504.html"
	VideoName = "怪奇物语 第四季"
)

// Escape 解析text转换为结构体
func Escape(videoText string) entiy.VideoInfo {
	str1 := strings.Replace(videoText, "\\/", "/", -1)
	str2 := strings.Split(str1, "=")[1]
	var videoInfo entiy.VideoInfo
	Unmarshal([]byte(str2), &videoInfo)
	return videoInfo
}

// GetHtml 请求url获取html页面元素
func GetHtml(url string) (io.ReadCloser, int) {
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Add("authority", "www.appmovie.cc")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("cookie", "cf_clearance=.Tvq6eV0MmSyI5CFFBgWwGlOBHN6ibuC.rAnM3hangc-1674872591-0-150; _gid=GA1.2.874431029.1674872594; _ga_4DJFNVEG0X=GS1.1.1674872593.69.1.1674873002.0.0.0; _ga=GA1.2.1732938208.1666775106; _gat_gtag_UA_145450513_1=1")
	req.Header.Add("referer", "https://www.appmovie.cc/index.php/vod/search.html?wd=%E6%80%AA%E5%A5%87%E7%89%A9%E8%AF%AD&submit=")
	req.Header.Add("sec-ch-ua", "\"Not_A Brand\";v=\"99\", \"Microsoft Edge\";v=\"109\", \"Chromium\";v=\"109\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 Edg/109.0.1518.69")
	res, err := client.Do(req)
	if err == nil {
		return res.Body, res.StatusCode
	}
	defer res.Body.Close()
	return nil, 403
}

// ParseHtml 解析请求html页面
func ParseHtml(body io.ReadCloser) string {
	var video entiy.Video
	var vts []entiy.VT
	var infoArray []string
	queryBig, _ := goquery.NewDocumentFromReader(body)
	videoName := queryBig.Find("div.stui-content__detail h3.title").Text()
	video.VideoName = videoName
	queryBig.Find("div.stui-content__detail p.data").Each(func(i int, selection *goquery.Selection) {
		infoArray = append(infoArray, strings.Join(strings.Fields(selection.Text()), ""))
	})
	video.Director = infoArray
	protagonist := queryBig.Find("div.stui-content__detail p.desc").Text()
	video.Protagonist = strings.Join(strings.Fields(protagonist), "")
	queryBig.Find("ul.stui-content__playlist li").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Find("a").Attr("href")
		text := VideoName + selection.Find("a").Text()
		fmt.Println(text, "- ", href, " - ", Url+href)
		html, code := GetHtml(Url + href)
		if code == http.StatusOK {
			querySmall, _ := goquery.NewDocumentFromReader(html)
			videoText := querySmall.Find("div.pl-box div.pl-l div.stui-player__video script").First().Text()
			videoInfo := Escape(videoText)
			vt := entiy.VT{
				FileName: text,
				Url:      videoInfo.Url,
				UrlNext:  videoInfo.UrlNext,
			}
			vts = append(vts, vt)
		}
	})
	video.Vts = vts
	marshal, _ := Marshal(video)
	return string(marshal)
}

// 主方法
func main() {
	html, code := GetHtml(VideoUrl)
	if code == http.StatusOK {
		parseHtml := ParseHtml(html)
		create, _ := os.Create(VideoName + ".json")
		create.WriteString(parseHtml)
	}
}
