## Comparaison gofeed / norme jsonfeed

|Variable|Type|Gofeed|Balise|O|Description|Note|
|-|-|-|-|-|-|-|
|ID|String|Y|Y|Y|Ideally, the id is the full URL of the resource described by the item, since URLs make great unique identifiers.|A Check|
|URL|String|Y|Y|N|URL of the resource described by the item. It’s the permalink|
|ExternalURL|String|Y|Y|N|external_url (very optional, string) is the URL of a page elsewhere. This is especially useful for linkblogs|Probablement Inutile pour flowwatcher|
|Title|String|Y|Y|N|Titre|Peut être non présent avec les micro blogs|
|ContentHTML|String|Y|Y| content_html and content_text are each optional strings — but one or both must be present.| |

O: Obligatoire


 
 
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
