package entiy

/**
 * @title dao
 * @author CH00SE1
 * @date 2022-12-16 15:22:42
 */

// VideoInfo 网页视频对象
type VideoInfo struct {
	Flag     string `json:"flag"`
	Encrypt  int    `json:"encrypt"`
	Trysee   int    `json:"trysee"`
	Points   int    `json:"points"`
	Link     string `json:"link"`
	LinkNext string `json:"link_next"`
	LinkPre  string `json:"link_pre"`
	Url      string `json:"url"`
	UrlNext  string `json:"url_next"`
	From     string `json:"from"`
	Server   string `json:"server"`
	Note     string `json:"note"`
}

// Video 剧集信息
type Video struct {
	VideoName   string   `json:"video_name"`
	Director    []string `json:"director"`
	Protagonist string   `json:"protagonist"`
	Vts         []VT     `json:"vts"`
}

// VT 视频内容信息
type VT struct {
	FileName string `json:"file_name"`
	Url      string `json:"url"`
	UrlNext  string `json:"url_next"`
}
