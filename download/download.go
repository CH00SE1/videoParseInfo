package main

import (
	m3u8 "github.com/CH00SE1/m3u8Download"
	"time"
)

/**
 * @title download
 * @author CH00SE1
 * @date 2022-12-15 11:58:04
 */

func main() {

	m3u8.VideoInfo{
		FileName: "康先生 “好爽~老公射给我!!” 大老板酒店约炮高挑女公关~呻吟到底就是不要在桌上玩~",

		Url: "https://cdn1.xjzyplay.com/20230125/NygDfHi1/1500kb/hls/index.m3u8",

		Output: "G:\\install\\mv\\" + time.Now().Format("2006-01-02") + "\\",

		ChanSize: 1 << 6,
	}.DownloadM3u8()

}

//FileName: "118job00038 働くオンナ2 VOL.42",

//Url: "https://video.jcwbf.com/20221217/OTM2NmQ2N2/102743/1280/hls/encrypt/index.m3u8",
