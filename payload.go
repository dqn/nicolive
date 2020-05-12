package nicolive

import "encoding/xml"

type getplayerstatus struct {
	XMLName xml.Name             `xml:"getplayerstatus"`
	Text    string               `xml:",chardata"`
	Status  string               `xml:"status,attr"`
	Time    string               `xml:"time,attr"`
	Stream  stream               `xml:"stream"`
	User    user                 `xml:"user"`
	Ms      ms                   `xml:"ms"`
	TidList string               `xml:"tid_list"`
	Tickets tickets              `xml:"tickets"`
	Twitter twitter              `xml:"twitter"`
	Player  player               `xml:"player"`
	Marquee marquee              `xml:"marquee"`
	Error   getplayerstatusError `xml:"error"`
}

type getplayerstatusError struct {
	Text string `xml:",chardata"`
	Code code   `xml:"code"`
}

type code struct {
	Text string `xml:",chardata"`
}

type stream struct {
	Text                         string       `xml:",chardata"`
	ID                           string       `xml:"id"`
	Title                        string       `xml:"title"`
	Description                  string       `xml:"description"`
	ProviderType                 string       `xml:"provider_type"`
	DefaultCommunity             string       `xml:"default_community"`
	International                string       `xml:"international"`
	IsOwner                      string       `xml:"is_owner"`
	OwnerID                      string       `xml:"owner_id"`
	OwnerName                    string       `xml:"owner_name"`
	IsReserved                   string       `xml:"is_reserved"`
	IsNiconicoEnqueteEnabled     string       `xml:"is_niconico_enquete_enabled"`
	WatchCount                   string       `xml:"watch_count"`
	CommentCount                 string       `xml:"comment_count"`
	BaseTime                     string       `xml:"base_time"`
	OpenTime                     string       `xml:"open_time"`
	StartTime                    string       `xml:"start_time"`
	EndTime                      string       `xml:"end_time"`
	IsRerunStream                string       `xml:"is_rerun_stream"`
	IsArchiveplayserver          string       `xml:"is_archiveplayserver"`
	BourbonURL                   string       `xml:"bourbon_url"`
	FullVideo                    string       `xml:"full_video"`
	AfterVideo                   string       `xml:"after_video"`
	BeforeVideo                  string       `xml:"before_video"`
	KickoutVideo                 string       `xml:"kickout_video"`
	TwitterTag                   string       `xml:"twitter_tag"`
	DanjoCommentMode             string       `xml:"danjo_comment_mode"`
	Aspect                       string       `xml:"aspect"`
	InfinityMode                 string       `xml:"infinity_mode"`
	Archive                      string       `xml:"archive"`
	Press                        press        `xml:"press"`
	PluginDelay                  string       `xml:"plugin_delay"`
	PluginURL                    string       `xml:"plugin_url"`
	PluginUrls                   string       `xml:"plugin_urls"`
	AllowNetduetto               string       `xml:"allow_netduetto"`
	NgScoring                    string       `xml:"ng_scoring"`
	IsNonarchiveTimeshiftEnabled string       `xml:"is_nonarchive_timeshift_enabled"`
	IsTimeshiftReserved          string       `xml:"is_timeshift_reserved"`
	HeaderComment                string       `xml:"header_comment"`
	FooterComment                string       `xml:"footer_comment"`
	SplitBottom                  string       `xml:"split_bottom"`
	SplitTop                     string       `xml:"split_top"`
	BackgroundComment            string       `xml:"background_comment"`
	FontScale                    string       `xml:"font_scale"`
	CommentLock                  string       `xml:"comment_lock"`
	Telop                        telop        `xml:"telop"`
	ContentsList                 contentsList `xml:"contents_list"`
	PictureURL                   string       `xml:"picture_url"`
	ThumbURL                     string       `xml:"thumb_url"`
	IsPriorityPrefecture         string       `xml:"is_priority_prefecture"`
}

type press struct {
	Text         string `xml:",chardata"`
	DisplayLines string `xml:"display_lines"`
	DisplayTime  string `xml:"display_time"`
	StyleConf    string `xml:"style_conf"`
}

type telop struct {
	Text   string `xml:",chardata"`
	Enable string `xml:"enable"`
}

type contentsList struct {
	Text     string   `xml:",chardata"`
	Contents contents `xml:"contents"`
}

type contents struct {
	Text         string `xml:",chardata"`
	ID           string `xml:"id,attr"`
	DisableAudio string `xml:"disableAudio,attr"`
	DisableVideo string `xml:"disableVideo,attr"`
	StartTime    string `xml:"start_time,attr"`
}

type user struct {
	Text           string      `xml:",chardata"`
	UserID         string      `xml:"user_id"`
	Nickname       string      `xml:"nickname"`
	IsPremium      string      `xml:"is_premium"`
	UserAge        string      `xml:"userAge"`
	UserSex        string      `xml:"userSex"`
	UserDomain     string      `xml:"userDomain"`
	UserPrefecture string      `xml:"userPrefecture"`
	UserLanguage   string      `xml:"userLanguage"`
	RoomLabel      string      `xml:"room_label"`
	RoomSeetno     string      `xml:"room_seetno"`
	IsJoin         string      `xml:"is_join"`
	TwitterInfo    twitterInfo `xml:"twitter_info"`
}

type twitterInfo struct {
	Text            string `xml:",chardata"`
	Status          string `xml:"status"`
	ScreenName      string `xml:"screen_name"`
	FollowersCount  string `xml:"followers_count"`
	IsVip           string `xml:"is_vip"`
	ProfileImageURL string `xml:"profile_image_url"`
	AfterAuth       string `xml:"after_auth"`
	TweetToken      string `xml:"tweet_token"`
}

type ms struct {
	Text   string `xml:",chardata"`
	Addr   string `xml:"addr"`
	Port   string `xml:"port"`
	Thread string `xml:"thread"`
}

type tickets struct {
	Text   string         `xml:",chardata"`
	Stream []ticketStream `xml:"stream"`
}

type ticketStream struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
}

type twitter struct {
	Text         string `xml:",chardata"`
	LiveEnabled  string `xml:"live_enabled"`
	VipModeCount string `xml:"vip_mode_count"`
	LiveApiURL   string `xml:"live_api_url"`
}

type player struct {
	Text                         string      `xml:",chardata"`
	QosAnalytics                 string      `xml:"qos_analytics"`
	DialogImage                  dialogImage `xml:"dialog_image"`
	IsNoticeViewerBalloonEnabled string      `xml:"is_notice_viewer_balloon_enabled"`
	ErrorReport                  string      `xml:"error_report"`
}

type dialogImage struct {
	Text    string `xml:",chardata"`
	Oidashi string `xml:"oidashi"`
}

type marquee struct {
	Text             string `xml:",chardata"`
	Category         string `xml:"category"`
	GameKey          string `xml:"game_key"`
	GameTime         string `xml:"game_time"`
	ForceNicowariOff string `xml:"force_nicowari_off"`
}
