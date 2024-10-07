## Comparaison gofeed / norme jsonfeed

|Variable|Type|Gofeed|Balise|O|Description|Note|
|-|-|-|-|-|-|-|
|ID|string|Y|Y|Y|Ideally, the id is the full URL of the resource described by the item, since URLs make great unique identifiers.|A Check|
|URL|string|Y|Y|N|URL of the resource described by the item. It’s the permalink|
|ExternalURL|string|Y|Y|N|external_url (very optional, string) is the URL of a page elsewhere. This is especially useful for linkblogs|Probablement Inutile pour flowwatcher|
|Title|string|Y|Y|N|Titre|Peut être non présent avec les micro blogs|
|ContentHTML|string|Y|Y|1|content_html and content_text are each optional strings — but one or both must be present.| |
|ContentText|string|Y|Y|1|Same as above ||
|Summary|string|Y|Y|N|summary (optional, string) is a plain text sentence or two describing the item.||
|Image|string|Y|Y|N|URL of the main image for the item. This image may also appear in the content_html||
|BannerImage|string |Y|Y|N| banner_image (optional, string) is the URL of an image to use as a banner.||
|DatePublished|string|Y|Y|N|specifies the date in RFC 3339 format. (Example: 2010-02-07T14:04:00-05:00.)|
|DateModified|string|Y|Y|N|specifies the modification date in RFC 3339 format.|Pas d'intérêt pour nous|
|Author|string|Y|Y|N| Has the same structure as the top-level author. If not specified in an item, then the top-level author, if present, is the author of the item.|
|Tags|[]string|can have any plain text values you want. Tags tend to be just one word, but they may be anything.||
|Attachments|*[]Attachments|Y|Y|N|lists related resources. Podcasts, for instance, would include an attachment that’s an audio or video file. An individual item may have one or more |Pas d'intérêt pour nous|
|Extensions|----|N|Y|N|Custom objects|Pas d'intérêt pour flo|

O: Obligatoire

<br>
<br>

En fonction des résultats...

En priorité pour:
- Le texte : Summary>Content HTML>Content Text
- Image	   : Image>Banner Image

