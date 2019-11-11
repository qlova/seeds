//Package a provides an html a element.
package a

import (
	"strings"
	"unsafe"

	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
)

//Seed is a a.
type Seed struct {
	seed.Seed
}

//New returns a new a.
func New(message ...string) Seed {
	var Text = seed.New()
	Text.SetTag("a")

	if len(message) > 0 {
		Text.SetText(message[0])
	}

	return Seed{Text}
}

//AddTo parent.
func AddTo(parent seed.Interface, message ...string) Seed {
	var Text = New(message...)
	parent.Root().Add(Text)
	return Text
}

/*
SetDownload prompts the user to save the linked URL instead of navigating to it.
Can be used with or without a value:
Without a value, the browser will suggest a filename/extension, generated from various sources:

* The Content-Disposition HTTP header
* The final segment in the URL path
* The media type (from the Content-Type header, the start of a data: URL, or Blob.type for a blob: URL)

Defining a value suggests it as the filename. / and \ characters are converted to underscores (_). Filesystems may forbid other characters in filenames, so browsers will adjust the suggested name if necessary.
*/
func (a Seed) SetDownload(download string) {
	a.Element.Set(html.Download, download)
}

/*
SetHypertextReference sets the URL that the hyperlink points to.
Links are not restricted to HTTP-based URLs — they can use any URL scheme supported by browsers:

* Sections of a page with fragment URLs
* Pieces of media files with media fragments
* Telephone numbers with tel: URLs
* Email addresses with mailto: URLs

While web browsers may not support other URL schemes, web sites can with registerProtocolHandler()
*/
func (a Seed) SetHypertextReference(url string) {
	a.Element.Set(html.HypertextReference, url)
}

/*
SetHypertextReferenceLanguage sets a hint at the human language of the linked URL.
No built-in functionality.
*/
func (a Seed) SetHypertextReferenceLanguage(language html.Language) {
	a.Element.Set(html.HypertextReferenceLanguage, string(language))
}

/*
SetPing sets a space-separated list of URLs.

When the link is followed, the browser will send
POST requests with the body PING to the URLs.
Typically for tracking.
*/
func (a Seed) SetPing(urls []string) {
	a.Element.Set(html.Ping, strings.Join(urls, " "))
}

//SetRelationship sets the relationship of the linked URL as space-separated link types.
func (a Seed) SetRelationship(linktypes []html.LinkType) {
	s := *(*[]string)(unsafe.Pointer(&linktypes))
	a.Element.Set(html.Relationship, strings.Join(s, " "))
}

/*
SetTarget sets where to display the linked URL,
as the name for a browsing context (a tab, window, or <iframe>).
*/
func (a Seed) SetTarget(target html.TargetType) {
	a.Element.Set(html.Target, string(target))
}

/*
SetType sets a hints at the linked URL’s format with a MIME type.
No built-in functionality.
*/
func (a Seed) SetType(mimetype string) {
	a.Element.Set(html.Type, mimetype)
}
