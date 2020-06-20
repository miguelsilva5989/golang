package main

import (
	"encoding/xml"
	"fmt"
)

type Location struct {
	Loc string `xml:"loc"`
}

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"` // slice to unmarshal the tag that is under
}

// to remove the {} in the array of Locations
func (l Location) String() string {
	return fmt.Sprint(l.Loc)
}

func main() {
	xmlstr := `
	<sitemapindex>
		<sitemap>
			<loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
		</sitemap>
		<sitemap>
			<loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
		</sitemap>
		<sitemap>
			<loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
		</sitemap>
	</sitemapindex>
	`
	// fmt.Println(xmlstr)

	var s SitemapIndex
	// https://golang.org/src/encoding/xml/example_test.go
	err := xml.Unmarshal([]byte(xmlstr), &s)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	for _, Location := range s.Locations {
		fmt.Println(Location)
	}

}
