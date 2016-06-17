package mashable

/*
import (
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func Parse(c *collection) string {
	resp, err := http.Get(c.URL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("got url", c.URL)
	log.Println("got response with", resp.Status)
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		root, err := html.Parse(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		content, ok := scrape.Find(root, scrape.ById("post-slider"))
		if !ok {
			log.Fatal("failed to find content")
		}

		article, ok := scrape.Find(content, scrape.ByClass("post"))
		if !ok {
			log.Fatal("failed to find article")
		}

		id := scrape.Attr(article, "data-id")
		log.Println("got id ", id)
		return id

	} else {
		log.Println("response contains error, exiting...")
		log.Fatal()
	}
	return ""
}
*/
