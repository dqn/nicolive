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
	fmt.Println("Time:", g.Time)
	fmt.Println("LiveID:", g.Stream.ID)
	fmt.Println("Title:", g.Stream.Title)
	fmt.Println("Description:", g.Stream.Description)
	fmt.Println("ProviderType:", g.Stream.ProviderType)
	fmt.Println("WatchCount:", g.Stream.WatchCount)
	fmt.Println("CommentCount:", g.Stream.CommentCount)
	fmt.Println("StartTime:", g.Stream.StartTime)
	fmt.Println("EndTime:", g.Stream.EndTime)
	fmt.Println("UserID:", g.User.UserID)
	fmt.Println("Nickname:", g.User.Nickname)
	fmt.Println("IsPremium:", g.User.IsPremium)
	fmt.Println("UserAge:", g.User.UserAge)
	fmt.Println("UserSex:", g.User.UserSex)
	fmt.Println("UserDomain:", g.User.UserDomain)
	fmt.Println("UserPrefecture:", g.User.UserPrefecture)
	fmt.Println("UserLanguage:", g.User.UserLanguage)
	fmt.Println("RoomLabel:", g.User.RoomLabel)
	fmt.Println("RoomSeetno:", g.User.RoomSeetno)
	fmt.Println("Addr:", g.Ms.Addr)
	fmt.Println("Port:", g.Ms.Port)
	fmt.Println("Thread:", g.Ms.Thread)
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
		ResFrom: "-1",
		Version: "20061206",
		Scores:  "1",
	}

	b, err := xml.Marshal(&m)
	if err != nil {
		return nil, err
	}
	b = append(b, 0)

	return b, nil
}

func (n *Nicolive) Listen(liveID string) error {
	g, err := n.getPlayerStatus(liveID)
	if err != nil {
		return err
	}

	printPlayerstatus(g)

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
	var c chat

	// Initially, thread and chat are combined.
	if i := bytes.Index(b, []byte("<chat")); i > 0 {
		if xml.Unmarshal(b[:i], &t) == nil {
			println(t.Thread)
		}
		b = b[i:]
	}

	for {
		err = xml.Unmarshal(b, &c)
		if err != nil {
			return err
		}

		println(string(b))

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
