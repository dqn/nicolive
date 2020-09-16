package nicolive

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const (
	resFrom = "-1"
	version = "20061206"
	scores  = "1"
)

type Nicolive struct {
	client *http.Client
}

func New(mail, password string) (*Nicolive, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Jar: jar}

	resp, err := client.PostForm(
		"https://account.nicovideo.jp/api/v1/login",
		url.Values{
			"mail_tel": {mail},
			"password": {password},
		},
	)
	if err != nil {
		return nil, err
	}

	if f, ok := resp.Header["X-Niconico-Authflag"]; !ok || f[0] != "1" {
		err = fmt.Errorf("failed to login")
		return nil, err
	}

	return &Nicolive{client}, nil
}

func printPlayerstatus(g *getplayerstatus) {
	fmt.Println("LiveID:", g.Stream.ID)
	fmt.Println("Title:", g.Stream.Title)
	fmt.Println("Description:", g.Stream.Description)
	fmt.Println("ProviderType:", g.Stream.ProviderType)
	fmt.Println("WatchCount:", g.Stream.WatchCount)
	fmt.Println("CommentCount:", g.Stream.CommentCount)
	fmt.Println("StartTime:", g.Stream.StartTime)
	fmt.Println("EndTime:", g.Stream.EndTime)
	fmt.Println("RoomLabel:", g.User.RoomLabel)
}

func connect(addr, port string) (net.Conn, error) {
	address := fmt.Sprintf("%s:%s", addr, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func makeMessage(thread string) ([]byte, error) {
	m := message{
		Thread:  thread,
		ResFrom: resFrom,
		Version: version,
		Scores:  scores,
	}

	b, err := xml.Marshal(&m)
	if err != nil {
		return nil, err
	}
	b = append(b, 0)

	return b, nil
}

func (n *Nicolive) Listen(liveID string, handler func(c *Chat) error) error {
	g, err := n.getPlayerStatus(liveID)
	if err != nil {
		return err
	}

	// process playerstatus?
	// printPlayerstatus(g)

	conn, err := connect(g.Ms.Addr, g.Ms.Port)
	if err != nil {
		return err
	}
	defer conn.Close()

	msg, err := makeMessage(g.Ms.Thread)
	if err != nil {
		return err
	}

	_, err = conn.Write(msg)
	if err != nil {
		return err
	}

	b := make([]byte, 1024)
	_, err = conn.Read(b)
	if err != nil {
		return err
	}

	var t thread
	var c Chat

	// Initially, thread and chat are combined.
	if i := bytes.Index(b, []byte("<chat")); i > 0 {
		if xml.Unmarshal(b[:i], &t) == nil {
			// process thread?
		}
		b = b[i:]
	}

	for {
		err = xml.Unmarshal(b, &c)
		if err != nil {
			return err
		}

		// println(string(b))
		err = handler(&c)
		if err != nil {
			return err
		}

		_, err = conn.Read(b)
		if err != nil {
			return err
		}
	}
}

func (n *Nicolive) getPlayerStatus(liveID string) (*getplayerstatus, error) {
	url := fmt.Sprintf("https://live.nicovideo.jp/api/getplayerstatus/%s", liveID)
	resp, err := n.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var g getplayerstatus
	err = xml.Unmarshal(b, &g)
	if err != nil {
		return nil, err
	}
	if g.Error.Code.Text != "" {
		err = fmt.Errorf("an error occurred: %s", g.Error.Code.Text)
		return nil, err
	}
	if g.Status != "ok" {
		err = fmt.Errorf("unknown error occurred:\n%s", string(b))
		return nil, err
	}

	return &g, nil
}
