package codes

type CodeInfo struct {
	LineOne   string
	LineTwo   string
	LineThree string
	LineFour  string
	LineFive  string
	LineSix   string
	LineSeven string
	LineEight string
}

type AllStatusCodes map[string]CodeInfo

var StatusCodes = AllStatusCodes{
	"100": {
		LineOne:   "Continue",
		LineTwo:   "This informational response code indicates",
		LineThree: "everything so far is ok and the client should",
		LineFour:  "continue with the request,",
		LineFive:  "or ignore it is it's already finished.",
		LineSix:   " ",
		LineSeven: " ",
		LineEight: " ",
	},
	"101": {
		LineOne:   "Switching Protocols",
		LineTwo:   "This informational response code indicates",
		LineThree: "a protocol to which the server switches.",
		LineFour:  "The protocol is specified in the Upgrade",
		LineFive:  "request header from the client.",
		LineSix:   " ",
		LineSeven: "Switching protocols might be used with WebSockets",
		LineEight: " ",
	},
	"102": {
		LineOne:   "Processing (WebDAV)",
		LineTwo:   "This informational response code is an interim",
		LineThree: "response to inform the client that the server",
		LineFour:  "has accepted the request but not yet completed it.",
		LineFive:  " ",
		LineSix:   "WebDAV - or Web Distributed Authoring and Versioning",
		LineSeven: "is an extension to HTTP that lets clients edit",
		LineEight: "remote content on the web.",
	},
	"103": {
		LineOne:   "Early Hints",
		LineTwo:   "This informational response code is primarily",
		LineThree: "intended to be used with the Link header",
		LineFour:  "to allow the user agent to start preloading",
		LineFive:  "resources while the server is still preparing a response.",
		LineSix:   " ",
		LineSeven: " ",
		LineEight: " ",
	},
	"200": {
		LineOne:   "OK",
		LineTwo:   "This successful response code indicates the request",
		LineThree: "has succeeded. A 200 response is cacheable by default.",
		LineFour:  "The meaning of a 200 response depends on the HTTP method:",
		LineFive:  "GET - Resource has been fetched and is in the message body",
		LineSix:   "HEAD - The representation headers are included in the response.",
		LineSeven: "POST - The resource describing the result of the action is in the message body",
		LineEight: "TRACE - The message body contains the request message as recieved by the server",
	},
	"201": {
		LineOne:   "Created",
		LineTwo:   "This successful response code indicates the request",
		LineThree: "has succeeded and led to the creation of a resource.",
		LineFour:  " ",
		LineFive:  "The common use case for this status code is as the result",
		LineSix:   "of a POST request.",
		LineSeven: " ",
		LineEight: " ",
	},
	"202": {
		LineOne:   "Accepted",
		LineTwo:   "This successful response code indicates the request",
		LineThree: "has been accepted for processing, but the processing has not been completed.",
		LineFour:  "In fact, the processing may not have started yet.",
		LineFive:  " ",
		LineSix:   "202 is non-committal, meaning there's no way for the server to later send",
		LineSeven: "an asynchronous response indicating the outcome of processing the request.",
		LineEight: " ",
	},
	"203": {
		LineOne:   "Non-Authoritative Information",
		LineTwo:   "This successful response code indicates that the request",
		LineThree: "was successful but the enclosed payload has been modified by",
		LineFour:  "a transforming proxy from that of the origin server's",
		LineFive:  "200 OK response.",
		LineSix:   " ",
		LineSeven: " ",
		LineEight: " ",
	},
	"204": {
		LineOne:   "No Content",
		LineTwo:   "This successful response code indicates that the request",
		LineThree: "has succeeded, but the client doesnt need to navigate away from its",
		LineFour:  "current page.",
		LineFive:  " ",
		LineSix:   "This can be used in an implementation of 'save and continue editing'",
		LineSeven: "where a PUT request is used to save the information and a 204 is the response",
		LineEight: "to indicate the editor should not be replaced by some other page.",
	},
	"205": {
		LineOne:   "Reset Content",
		LineTwo:   " ",
		LineThree: " ",
		LineFour:  " ",
		LineFive:  " ",
		LineSix:   " ",
		LineSeven: " ",
		LineEight: " ",
	},
	"206": {
		LineOne:   "Partial Content",
		LineTwo:   "This successful response code indicates that the request",
		LineThree: "has succeeded and the body contains the requested ranges of data",
		LineFour:  "as described in the Range header of the request.",
		LineFive:  " ",
		LineSix:   "If there is only one range, the Content-Type of the whole response is",
		LineSeven: "set to the type of the document and a Content-Range is provided.",
		LineEight: " ",
	},
	"207": {
		LineOne:   "Multi-Status (WebDAV)",
		LineTwo:   "This successful response code conveys information about multiple resources",
		LineThree: "for situations where multiple status codes might be appropriate.",
		LineFour:  " ",
		LineFive:  " ",
		LineSix:   "WebDAV - or Web Distributed Authoring and Versioning",
		LineSeven: "is an extension to HTTP that lets clients edit",
		LineEight: "remote content on the web.",
	},
	"208": {
		LineOne:   "Already Reported",
		LineTwo:   "This successful response code is used inside a <dav:propstate> response",
		LineThree: "element to avoid repeatedly enumerating the internal members of",
		LineFour:  "multiple bindings to the same collection.",
		LineFive:  " ",
		LineSix:   "WebDAV - or Web Distributed Authoring and Versioning",
		LineSeven: "is an extension to HTTP that lets clients edit",
		LineEight: "remote content on the web.",
	},
	"226": {
		LineOne:   "IM Used (HTTP Delta encoding)",
		LineTwo:   "This successful response code indicates the server has fulfilled a GET",
		LineThree: "request for the resource, and the response is a representation of",
		LineFour:  "the result of one or more instance-manipulations applied",
		LineFive:  "to the current instance.",
		LineSix:   " ",
		LineSeven: "https://datatracker.ietf.org/doc/html/rfc3229",
		LineEight: " ",
	},
	"300": {
		LineOne:   "Multiple Choices",
		LineTwo:   "This redirection message response code indicates that the request",
		LineThree: "has more than one possible responses.",
		LineFour:  "The user-agent or the user should choose one of them.",
		LineFive:  "As there's no standard way of choosing a response, this code is rarely used.",
		LineSix:   " ",
		LineSeven: "If the server has a preferred use, it should generate a Location header.",
		LineEight: " ",
	},
	"301": {
		LineOne:   "Moved Permanently",
		LineTwo:   "This redirection message response code indicates that the requested",
		LineThree: "resource has been definitively moved to the URL given by the Location headers.",
		LineFour:  " ",
		LineFive:  "A browser redirects to the new URL and search engines update their links.",
		LineSix:   " ",
		LineSeven: " ",
		LineEight: " ",
	},
	"302": {
		LineOne:   "Found",
		LineTwo:   "This redirection message response code indicates that the resource requested",
		LineThree: "has been temporarily moved to the URL given by the Location header.",
		LineFour:  " ",
		LineFive:  "A browser redirects to this new URL but search engines don't",
		LineSix:   "update their links.",
		LineSeven: " ",
		LineEight: " ",
	},
	"303": {
		LineOne:   "See Other",
		LineTwo:   "This redirection message response code indicates that the redirects don't",
		LineThree: "link to the requested resource itself, but to another page.",
		LineFour:  " ",
		LineFive:  "An example of 'another page' could be a confirmation page",
		LineSix:   " ",
		LineSeven: " ",
		LineEight: " ",
	},
	"304": {
		LineOne:   "Not Modified",
		LineTwo:   "This redirection message response code indicates that there is no need to",
		LineThree: "retransmit the requested resources. It is an implicit redirection to a",
		LineFour:  "cached resource. This happens when the method is a safe method,",
		LineFive:  "such as GET or HEAD, or when the request is conditional and uses an",
		LineSix:   "If-None-Match or If-Modified-Since header.",
		LineSeven: " ",
		LineEight: " ",
	},
	"305": {
		LineOne:   "Use Proxy",
		LineTwo:   "Defined in a previous version of the HTTP specification to indicate",
		LineThree: "that a requested response must be accessed by a proxy.",
		LineFour:  " ",
		LineFive:  "This has been deprecated due to security concerns.",
		LineSix:   " ",
		LineSeven: " ",
		LineEight: " ",
	},
	"306": {
		LineOne:   "unused",
		LineTwo:   "This response code is no longer used; its just reserved. It was used",
		LineThree: "in a previous version of the HTTP/1.1 specification.",
		LineFour:  " ",
		LineFive:  " ",
		LineSix:   " ",
		LineSeven: " ",
		LineEight: " ",
	},
	"307": {
		LineOne:   "Temporary Redirect",
		LineTwo:   "This redirection message response code indicates that the resource",
		LineThree: "requested has been temporarily moved to the URL given in the Location header.",
		LineFour:  " ",
		LineFive:  "The difference between this and 302 - Found is that 307 guarantees that",
		LineSix:   "the method and the body will not be changed when the redirect",
		LineSeven: "request is made. With 302 some old clients were incorrectly changing",
		LineEight: "the method to GET.",
	},
	"308": {
		LineOne:   "Permanent Redirect",
		LineTwo:   "This redirection message response code indicates that the resource",
		LineThree: "requested has been definitively moved to the URL given by the Location",
		LineFour:  "header. A browser redirects to the page and search engines update",
		LineFive:  "their links to the resource.",
		LineSix:   " ",
		LineSeven: "The request body will not be altered, whereas 301 - Moved Permanently",
		LineEight: "may incorrectly sometimes be changed to a GET method.",
	},
	"400": {
		LineOne:   "Bad Request",
		LineTwo:   "This client error response code indicates that the server cannot or",
		LineThree: "will not process the request due to something that is perceived to",
		LineFour:  "be a client error.",
		LineFive:  "e.g. a malformed request syntax, invalid request message framing, etc",
		LineSix:   " ",
		LineSeven: "WARNING: The client should not repeat this request without modification.",
		LineEight: " ",
	},
	"401": {
		LineOne:   "Unauthorized",
		LineTwo:   "This client error response code indicates that the client reequest",
		LineThree: "has not been completed because it lacks valid authentication",
		LineFour:  "credentials for the requested resource.",
		LineFive:  " ",
		LineSix:   "This is similar to 403 - Forbidden, except that in situations",
		LineSeven: "resulting in this status code, user authentication can allow access.",
		LineEight: " ",
	},
	"402": {
		LineOne:   "Payment Required",
		LineTwo:   "This client error response code is a nonstandard, experimental response",
		LineThree: "status code that is reserved for future use. ",
		LineFour:  " ",
		LineFive:  "This status code was created to enable digital cash or (micro) payment",
		LineSix:   "systems and would indicate that the requested content is not available",
		LineSeven: "until the client makes a payment.",
		LineEight: " ",
	},
	"403": {
		LineOne:   "Forbidden",
		LineTwo:   "This client error response code indicates that the server understands",
		LineThree: "the request but refuses to authorize it.",
		LineFour:  " ",
		LineFive:  "This status code is similar to 401 - Unauthorized but for 403,",
		LineSix:   "re-authorizing makes no difference. The access is permanently forbidden",
		LineSeven: "and tied to the application logic, such as insufficient rights to a resource.",
		LineEight: " ",
	},
	"404": {
		LineOne:   "Not Found",
		LineTwo:   "This client error response code indicates that the server cannot find the",
		LineThree: "requested resource. Links that lead to a 404 page are often called broken",
		LineFour:  "or dead links and can be subject to 'link rot'",
		LineFive:  " ",
		LineSix:   "A 404 response only indicates that the resource is missing; not whether",
		LineSeven: "the absence is temporary or permanent. If a resource is permanently removed",
		LineEight: "use the 410 - Gone status instead.",
	},
	"405": {
		LineOne:   "Method Not Allowed",
		LineTwo:   "This client error response code indicates that the server knows the",
		LineThree: "request method, but the target resource doesnt support this method.",
		LineFour:  " ",
		LineFive:  "The server MUST generate an Allow header field in a 405 status",
		LineSix:   "code response. The field must contain a list of methods that the",
		LineSeven: "target resource currently supports.",
		LineEight: " ",
	},
	"406": {
		LineOne:   "Not Acceptable",
		LineTwo:   "This client error response code indicates that the server cannot",
		LineThree: "produce a response matching the list of acceptable values defined",
		LineFour:  "in the requests proactive content negotitation headers, and the ",
		LineFive:  "server is unwilling to supply a default representation.",
		LineSix:   " ",
		LineSeven: "Proactive content negotiation headers include:",
		LineEight: "Accept, Accept-Encoding, Accept-Language",
	},
	"407": {
		LineOne:   "Proxy Authentication Required",
		LineTwo:   "This client error response code indicates that the request has not",
		LineThree: "been applied because it lacks valid authentication credentials for a",
		LineFour:  "proxy server that is between the browser and the server than can",
		LineFive:  "access the requested resource.",
		LineSix:   " ",
		LineSeven: "This status is sent with a Proxy-Authenticate header that contains",
		LineEight: "information on how to authorize correctly.",
	},
	"408": {
		LineOne:   "Request Timeout",
		LineTwo:   "This client error response code means that the server would like to",
		LineThree: "shut down this unused connection. It is sent on an idle connection",
		LineFour:  "by some servers, even without any previous request by the client.",
		LineFive:  " ",
		LineSix:   "A server should send the 'close' Connection header field in the response.",
		LineSeven: "This response is used much more since some browsers use HTTP pre-connection",
		LineEight: "mechanisms to speed up surfing.",
	},
	"409": {
		LineOne:   "Conflict",
		LineTwo:   "This client error response code indicates a request conflict with the current",
		LineThree: "state of the target resource.",
		LineFour:  " ",
		LineFive:  "Conflicts are most likely to occur in response to a PUT request.",
		LineSix:   "For example, you may get a 409 response when uploading a file that is older",
		LineSeven: "than the existing one on the server, resulting in a version control conflict.",
		LineEight: " ",
	},
	"410": {
		LineOne:   "Gone",
		LineTwo:   "This client error response code indicates that access to the target resource",
		LineThree: "is no longer available at the origin server and that this condition is",
		LineFour:  "likely to be permanent.",
		LineFive:  " ",
		LineSix:   "If you dont know whether this condition is temporary or permanent, a 404",
		LineSeven: "status code should be used instead.",
		LineEight: " ",
	},
	"411": {
		LineOne:   "Length Required",
		LineTwo:   "This client error response code indicates that the server refuses to accept",
		LineThree: "the request without a defined Content-Length header.",
		LineFour:  " ",
		LineFive:  " ",
		LineSix:   " ",
		LineSeven: " ",
		LineEight: " ",
	},
	"412": {
		LineOne:   "Precondition Failed",
		LineTwo:   "This client error response code indicates that access to the target resource",
		LineThree: "has been denied. This happens with conditional requests on methods other than",
		LineFour:  "GET or HEAD when the condition defined by the If-Unmodified-Since or ",
		LineFive:  "If-None-Match headers is not fulfilled. In that case, the request, ",
		LineSix:   "usually an upload or a modification of a resource, cannot be made",
		LineSeven: "and this error response is sent back.",
		LineEight: " ",
	},
	"413": {
		LineOne:   "Payload Too Large",
		LineTwo:   "This client error response code indicates that the request entity",
		LineThree: "is larger than limits defined by server; the server might close",
		LineFour:  "the connection or return a Retry-After header field.",
		LineFive:  " ",
		LineSix:   " ",
		LineSeven: " ",
		LineEight: " ",
	},
	"414": {
		LineOne:   "URI Too Long",
		LineTwo:   "This client error response code indicates that the URI requested",
		LineThree: "by the client is longer than the server is willing to interpret.",
		LineFour:  "There are a few rare conditions when this might occur:",
		LineFive:  "- When a client has improperly converted a POST to a GET request.",
		LineSix:   "- When the client is in a loop of redirection.",
		LineSeven: "- When the server is under attack by a client attempting an exploit.",
		LineEight: " ",
	},
	"415": {
		LineOne:   "Unsupported Media Type",
		LineTwo:   "This client error response code indicates that the server refuses",
		LineThree: "to accept the request because the payload format is in an ",
		LineFour:  "unsupported format.",
		LineFive:  " ",
		LineSix:   "The format problem might be due to the requests indicated Content-Type",
		LineSeven: "or Content-Encoding, or as a result of inspecting the data directly.",
		LineEight: " ",
	},
	"416": {
		LineOne:   "Range Not Satisfiable",
		LineTwo:   "This client error response code indicates that a server cannot serve",
		LineThree: "the requested ranges. The most likely reason is that the document",
		LineFour:  "doesnt contain such ranges, or that the Range header value, though",
		LineFive:  "syntactically correct, doesnt make sense.",
		LineSix:   " ",
		LineSeven: "The 416 response message contains a Content-Range indicating an",
		LineEight: "unsatisfied range followed by a '/' and the current length of the resource.",
	},
	"417": {
		LineOne:   "Expectation Failed",
		LineTwo:   "This client error response code indicates that the expectation given",
		LineThree: "in the requests Expect header could not be met.",
		LineFour:  " ",
		LineFive:  " ",
		LineSix:   " ",
		LineSeven: " ",
		LineEight: " ",
	},
	"418": {
		LineOne:   "I'm A Teapot",
		LineTwo:   "                     ;,'",
		LineThree: "           _o_      ;:;'",
		LineFour:  "   ,-.'---'.__ ;",
		LineFive:  " ((j;=====',-'",
		LineSix:   "'-\\     /",
		LineSeven: "      '-=-'     ",
		LineEight: "You should Google this one...",
	},
	"421": {
		LineOne:   "Misdirected Request",
		LineTwo:   "This client error response code indicates the request was directed at",
		LineThree: "a server that is not able to produce a response. This can be sent",
		LineFour:  "by a server that is not configured to produce responses for the",
		LineFive:  "combination of scheme and authority that are included in the ",
		LineSix:   "request URI",
		LineSeven: " ",
		LineEight: " ",
	},
	"422": {
		LineOne:   "Unprocessable Entity (WebDAV)",
		LineTwo:   "This client error response code indicates that the request was",
		LineThree: "well-formed but was unable to be followed due to semantic errors.",
		LineFour:  " ",
		LineFive:  " ",
		LineSix:   "WebDAV - or Web Distributed Authoring and Versioning",
		LineSeven: "is an extension to HTTP that lets clients edit",
		LineEight: "remote content on the web.",
	},
	"423": {
		LineOne:   "Locked (WebDAV)",
		LineTwo:   "This client error response code indicates that the resource",
		LineThree: "being accessed is locked.",
		LineFour:  " ",
		LineFive:  " ",
		LineSix:   "WebDAV - or Web Distributed Authoring and Versioning",
		LineSeven: "is an extension to HTTP that lets clients edit",
		LineEight: "remote content on the web.",
	},
	"424": {
		LineOne:   "Failed Dependency (WebDAV)",
		LineTwo:   "This client error response code indicates that the request",
		LineThree: "failed due to failure of a previous request.",
		LineFour:  " ",
		LineFive:  " ",
		LineSix:   "WebDAV - or Web Distributed Authoring and Versioning",
		LineSeven: "is an extension to HTTP that lets clients edit",
		LineEight: "remote content on the web.",
	},
	"425": {
		LineOne:   "Too Early",
		LineTwo:   "This client error response code indicates that the server",
		LineThree: "is unwilling to risk processing a request that might be replayed",
		LineFour:  "which creates the potential for a replay attack.",
		LineFive:  " ",
		LineSix:   " ",
		LineSeven: " ",
		LineEight: " ",
	},
	"426": {
		LineOne:   "Upgrade Required",
		LineTwo:   "This client error response code indicates that the server refuses",
		LineThree: "to perform the request using the current protocol but might be willing",
		LineFour:  "to do so after the client upgrades to a different protocol.",
		LineFive:  " ",
		LineSix:   "The server sends an Upgrade header with this response to indicate",
		LineSeven: "the required protocol(s)",
		LineEight: " ",
	},
	"428": {
		LineOne:   "Precondition Required",
		LineTwo:   "This client error response code indicates that the server requires",
		LineThree: "the request to be conditional.",
		LineFour:  " ",
		LineFive:  "Typically this means that a required precondition header such as If-Match",
		LineSix:   "is missing from the request.",
		LineSeven: " ",
		LineEight: " ",
	},
	"429": {
		LineOne:   "Too Many Requests",
		LineTwo:   "This client error response code indicates the user has sent too many",
		LineThree: "requests in a given amount of time ('rate limiting')",
		LineFour:  " ",
		LineFive:  "A Retry-After header might be included to this response indicating",
		LineSix:   "how long to wait before making a new request.",
		LineSeven: " ",
		LineEight: " ",
	},
	"431": {
		LineOne:   "Request Header Fields Too Large",
		LineTwo:   "This client error response code indicates that the server refuses",
		LineThree: "to process the request because the requests headers are too long.",
		LineFour:  " ",
		LineFive:  "Servers will often produce this status if:",
		LineSix:   "- The Referer URL is too long.",
		LineSeven: "- There are too many Cookies sent in the request.",
		LineEight: " ",
	},
	"451": {
		LineOne:   "Unavailable For Legal Reasons",
		LineTwo:   "This client error response code indicates that the user requested",
		LineThree: "a resource that is not available due to legal reasons, such as a",
		LineFour:  "web page for which a legal action has been issued.",
		LineFive:  " ",
		LineSix:   " ",
		LineSeven: " ",
		LineEight: " ",
	},
	"500": {
		LineOne:   "Internal Server Error",
		LineTwo:   "This server error response code indicates that the server encountered",
		LineThree: "an unexpected condition that prevented it from fulfilling the request.",
		LineFour:  " ",
		LineFive:  "This error is a generic catch-all response. Usually, this indicates",
		LineSix:   "the server cannot find a better 5xx response.",
		LineSeven: " ",
		LineEight: " ",
	},
	"501": {
		LineOne:   "Not Implemented",
		LineTwo:   "This server error response code indicates that the server does not support",
		LineThree: "the functionality required to fulfill the request.",
		LineFour:  " ",
		LineFive:  "This is the appropriate response when the server does not recognize the",
		LineSix:   "request method and is incapable of supporting it for any resource.",
		LineSeven: "The only methods that servers are required to support (and therefore",
		LineEight: "muust not return 501) are GET and HEAD",
	},
	"502": {
		LineOne:   "Bad Gateway",
		LineTwo:   "This server error response indicates that the server, while acting as",
		LineThree: "a gateway or proxy, received an invalid response from the upstream server.",
		LineFour:  " ",
		LineFive:  "Note: a gateway might refer to different things in networking, and a 502",
		LineSix:   "error is usually not something you can fix, but requires a fix on the ",
		LineSeven: "server or the proxies you are trying to get access through.",
		LineEight: " ",
	},
	"503": {
		LineOne:   "Service Unavailable",
		LineTwo:   "This server error response indicates that the server is not ready to",
		LineThree: "handle the request.",
		LineFour:  " ",
		LineFive:  "Common causes are a server that is down for maintenance or that is overloaded.",
		LineSix:   "This response should be used for temporary conditions and the Retry-After ",
		LineSeven: "header should, if possible, contain the estimated time for recovery.",
		LineEight: " ",
	},
	"504": {
		LineOne:   "Gateway Timeout",
		LineTwo:   "This server error response indicates that the server, while acting as a",
		LineThree: "gateway or proxy, did not get a response in time from the upstream",
		LineFour:  "server that it needed in order to complete the request.",
		LineFive:  " ",
		LineSix:   "Note: a gateway might refer to different things in networking, and a 502",
		LineSeven: "error is usually not something you can fix, but requires a fix on the",
		LineEight: "server or the proxies you are trying to get access through.",
	},
	"505": {
		LineOne:   "HTTP Version Not Supported",
		LineTwo:   "This server error response indicates that the HTTP version used",
		LineThree: "in the request is not supported by the server.",
		LineFour:  " ",
		LineFive:  " ",
		LineSix:   " ",
		LineSeven: " ",
		LineEight: " ",
	},
	"506": {
		LineOne:   "Variant Also Negotiates",
		LineTwo:   "This server error response indicates an internal server configuration",
		LineThree: "error in which the chosen variant is itself configured to engage",
		LineFour:  "in content negotiation, so is not a proper negotiation endpoint.",
		LineFive:  " ",
		LineSix:   " ",
		LineSeven: " ",
		LineEight: " ",
	},
	"507": {
		LineOne:   "Insufficient Storage",
		LineTwo:   "This server error response indicates that a method could not be performed",
		LineThree: "because the server cannot store the representation needed to ",
		LineFour:  "successfully complete the request.",
		LineFive:  " ",
		LineSix:   " ",
		LineSeven: " ",
		LineEight: " ",
	},
	"508": {
		LineOne:   "Loop Detected",
		LineTwo:   "This server error response indicates that the server terminated an",
		LineThree: "operation because it encountered an infinite loop while processing",
		LineFour:  "a request with 'Depth: infinity'. This status indicates that",
		LineFive:  "the entire operation failed.",
		LineSix:   " ",
		LineSeven: " ",
		LineEight: " ",
	},
	"510": {
		LineOne:   "Not Extended",
		LineTwo:   "This server error response code is sent in the context of the",
		LineThree: "HTTP Extension Framework, defined in RFC 2774.",
		LineFour:  " ",
		LineFive:  "In that specification a client may send a request that contains an",
		LineSix:   "extension to be used. If the server receives such a request, but any",
		LineSeven: "described extensions are not supported for the request, then the",
		LineEight: "server response with a 510.",
	},
	"511": {
		LineOne:   "Network Authentication Required",
		LineTwo:   "This server error response indicates that the client needs to",
		LineThree: "authenticate to gain network access.",
		LineFour:  " ",
		LineFive:  "The status is not generated by origin servers, but by intercepting",
		LineSix:   "proxies that control access to the network.",
		LineSeven: " ",
		LineEight: " ",
	},
}
