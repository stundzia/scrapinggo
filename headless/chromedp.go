package headless

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"github.com/stundzia/scrapinggo/fetch"
)

func ChromeDPExample() {
	opts := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
		chromedp.ProxyServer("pr.oxylabs.io:7777"),
		chromedp.Flag("start-fullscreen", true),
	}

	allocContext, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocContext)

	defer cancel()

	var res string
	var resUrl string
	var title string

	start := time.Now()
	err := chromedp.Run(
		ctx,
		network.Enable(),
		network.SetExtraHTTPHeaders(fetch.HeadersDesktopMap),
		chromedp.Navigate(fetch.TargetGoogleSearch),
		chromedp.Location(&resUrl),
		chromedp.Sleep(2*time.Second),
		chromedp.Navigate(fetch.TargetIPInfoJSON),
		chromedp.Sleep(3*time.Second),
		chromedp.OuterHTML("html", &res),
		chromedp.Navigate("https://github.com/"),
		chromedp.Sleep(1*time.Second),
		chromedp.Title(&title),
		chromedp.Click("input[type='text']"),
		chromedp.SendKeys("input[type='text']", "scrapinggo"),
		chromedp.Sleep(2*time.Second),
		chromedp.KeyEvent(kb.Enter),
		chromedp.Sleep(3*time.Second),
		chromedp.Click("div.f4.text-normal"),
		chromedp.Sleep(4*time.Second),
		chromedp.Navigate("http://bash.org"),
		chromedp.Sleep(time.Second),
	)
	if err != nil {
		fmt.Println("Err: ", err)
	}
	fmt.Println("Done in: ", time.Now().Sub(start))
	fmt.Println("IP Info Content: ", res)
	fmt.Println("Github Title: ", title)

	time.Sleep(2 * time.Second)
}
