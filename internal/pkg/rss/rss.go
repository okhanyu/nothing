package rss

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/feeds"
	"log"
	"nothing/internal/app/blog/model/post"
	"nothing/internal/app/blog/model/setting"
	"time"
)

type SettingRss struct {
	Title       string
	Link        string
	Description string
	Author      *feeds.Author
}

func Rss(settings *setting.SettingBo, boPost []*post.SimplePostVo) string {
	// 创建一个新的RSS Feed
	var settingRss SettingRss
	err := json.Unmarshal([]byte(settings.Config), &settingRss)
	if err != nil {
		return ""
	}
	feed := &feeds.Feed{
		Title:       settingRss.Title,
		Link:        &feeds.Link{Href: settingRss.Link},
		Description: settingRss.Description,
		Created:     time.Now(),
		Author:      settingRss.Author,
	}

	for _, bo := range boPost {
		// 添加一些示例项
		item := &feeds.Item{
			Title:       bo.Title,
			Link:        &feeds.Link{Href: fmt.Sprintf("https://hanyu.life/blog/detail.html?id=%d", bo.ID)},
			Description: bo.Summary,
			Created:     *bo.CreatedAt,
			Author:      settingRss.Author,
		}
		feed.Add(item)
	}

	// 生成RSS XML
	rss, err := feed.ToRss()
	if err != nil {
		log.Println(err)
	}

	return rss
}
