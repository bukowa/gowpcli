package fetch

//Attempts to convert a URL into embed HTML.

type Fetch struct {
    
    Url string // <url>
    
    Width string // [--width=<width>]
    
    Height string // [--height=<height>]
    
    PostId string // [--post-id=<id>]
    
    Discover bool // [--discover]
    
    SkipCache bool // [--skip-cache]
    
    SkipSanitization bool // [--skip-sanitization]
    
    DoShortcode bool // [--do-shortcode]
    
    LimitResponseSize string // [--limit-response-size=<size>]
    
    Raw bool // [--raw]
    
    RawFormat string // [--raw-format=<json|xml>]
    
}

//In non-raw mode, starts by checking the URL against the regex of the registered embed handlers.
//If none of the regex matches and it's enabled, then the URL will be given to the WP_oEmbed class.
//
//In raw mode, checks the providers directly and returns the data.
//
//## OPTIONS
//
//<url>
//: URL to retrieve oEmbed data for.
//
//[--width=<width>]
//: Width of the embed in pixels.
//
//[--height=<height>]
//: Height of the embed in pixels.
//
//[--post-id=<id>]
//: Cache oEmbed response for a given post.
//
//[--discover]
//: Enable oEmbed discovery. Defaults to true.
//
//[--skip-cache]
//: Ignore already cached oEmbed responses. Has no effect if using the 'raw' option, which doesn't use the cache.
//
//[--skip-sanitization]
//: Remove the filter that WordPress from 4.4 onwards uses to sanitize oEmbed responses. Has no effect if using the 'raw' option, which by-passes sanitization.
//
//[--do-shortcode]
//: If the URL is handled by a registered embed handler and returns a shortcode, do shortcode and return result. Has no effect if using the 'raw' option, which by-passes handlers.
//
//[--limit-response-size=<size>]
//: Limit the size of the resulting HTML when using discovery. Default 150 KB (the standard WordPress limit). Not compatible with 'no-discover'.
//
//[--raw]
//: Return the raw oEmbed response instead of the resulting HTML. Ignores the cache and does not sanitize responses or use registered embed handlers.
//
//[--raw-format=<json|xml>]
//: Render raw oEmbed data in a particular format. Defaults to json. Can only be specified in conjunction with the 'raw' option.
//---
//options:
//  - json
//  - xml
//---
//
//## EXAMPLES
//
//    # Get embed HTML for a given URL.
//    $ wp embed fetch https://www.youtube.com/watch?v=dQw4w9WgXcQ
//    <iframe width="525" height="295" src="https://www.youtube.com/embed/dQw4w9WgXcQ?feature=oembed" ...
//
//    # Get raw oEmbed data for a given URL.
//    $ wp embed fetch https://www.youtube.com/watch?v=dQw4w9WgXcQ --raw
//    {"author_url":"https:\/\/www.youtube.com\/user\/RickAstleyVEVO","width":525,"version":"1.0", ...
//
//