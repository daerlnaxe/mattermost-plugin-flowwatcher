## Comparaison gofeed / norme jsonfeed

|Variable|Gofeed|Balise|Requis|Type|Unique|Description|
|-|-|-|-|-|-|-|
|ID|Y|Y|Y|String|Y| Ideally, the id is the full URL of the resource described by the item, since URLs make great unique identifiers.|



	URL           string  `json:"url,omitempty"`            // url (optional, string) is the URL of the resource described by the item. It’s the permalink
	ExternalURL   string  `json:"external_url,omitempty"`   // external_url (very optional, string) is the URL of a page elsewhere. This is especially useful for linkblogs
	Title         string  `json:"title,omitempty"`          // title (optional, string) is plain text. Microblog items in particular may omit titles.
	ContentHTML   string  `json:"content_html,omitempty"`   // content_html and content_text are each optional strings — but one or both must be present. This is the HTML or plain text of the item. Important: the only place HTML is allowed in this format is in content_html. A Twitter-like service might use content_text, while a blog might use content_html. Use whichever makes sense for your resource. (It doesn’t even have to be the same for each item in a feed.)
	ContentText   string  `json:"content_text,omitempty"`   // Same as above
	Summary       string  `json:"summary,omitempty"`        // summary (optional, string) is a plain text sentence or two describing the item.
	Image         string  `json:"image,omitempty"`          // image (optional, string) is the URL of the main image for the item. This image may also appear in the content_html
	BannerImage   string  `json:"banner_image,omitempty"`   // banner_image (optional, string) is the URL of an image to use as a banner.
	DatePublished string  `json:"date_published,omitempty"` // date_published (optional, string) specifies the date in RFC 3339 format. (Example: 2010-02-07T14:04:00-05:00.)
	DateModified  string  `json:"date_modified,omitempty"`  // date_modified (optional, string) specifies the modification date in RFC 3339 format.
	Author        *Author `json:"author,omitempty"`         // author (optional, object) has the same structure as the top-level author. If not specified in an item, then the top-level author, if present, is the author of the item.

	Tags        []string       `json:"tags,omitempty"`        // tags (optional, array of strings) can have any plain text values you want. Tags tend to be just one word, but they may be anything.
	Attachments *[]Attachments `json:"attachments,omitempty"` // attachments (optional, array) lists related resources. Podcasts, for instance, would include an attachment that’s an audio or video file. An individual item may have one or more attachments.
	// TODO Extensions

	// Version 1.1
	Authors  []*Author `json:"authors,omitempty"`
	Language string    `json:"language,omitempty"`
