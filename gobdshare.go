package gobdshare

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
)

func getDSELatestPrices()  {
	resp, err := soup.Get("https://www.dsebd.org/latest_share_price_scroll_l.php")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "class", "table-responsive inner-scroll").FindAll("tr")
	for _, link := range links {
		fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	}
}

