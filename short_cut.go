package boom

var (
	Continue           = wrapFunc(100) // RFC 7231, 6.2.1
	SwitchingProtocols = wrapFunc(101) // RFC 7231, 6.2.2
	Processing         = wrapFunc(102) // RFC 2518, 10.1
	EarlyHints         = wrapFunc(103) // RFC 8297

	OK                   = wrapFunc(200) // RFC 7231, 6.3.1
	Created              = wrapFunc(201) // RFC 7231, 6.3.2
	Accepted             = wrapFunc(202) // RFC 7231, 6.3.3
	NonAuthoritativeInfo = wrapFunc(203) // RFC 7231, 6.3.4
	NoContent            = wrapFunc(204) // RFC 7231, 6.3.5
	ResetContent         = wrapFunc(205) // RFC 7231, 6.3.6
	PartialContent       = wrapFunc(206) // RFC 7233, 4.1
	Multi                = wrapFunc(207) // RFC 4918, 11.1
	AlreadyReported      = wrapFunc(208) // RFC 5842, 7.1
	IMUsed               = wrapFunc(226) // RFC 3229, 10.4.1

	MultipleChoices   = wrapFunc(300) // RFC 7231, 6.4.1
	MovedPermanently  = wrapFunc(301) // RFC 7231, 6.4.2
	Found             = wrapFunc(302) // RFC 7231, 6.4.3
	SeeOther          = wrapFunc(303) // RFC 7231, 6.4.4
	NotModified       = wrapFunc(304) // RFC 7232, 4.1
	UseProxy          = wrapFunc(305) // RFC 7231, 6.4.5
	_                 = wrapFunc(306) // RFC 7231, 6.4.6 (Unused)
	TemporaryRedirect = wrapFunc(307) // RFC 7231, 6.4.7
	PermanentRedirect = wrapFunc(308) // RFC 7538, 3

	BadRequest                   = wrapFunc(400) // RFC 7231, 6.5.1
	Unauthorized                 = wrapFunc(401) // RFC 7235, 3.1
	PaymentRequired              = wrapFunc(402) // RFC 7231, 6.5.2
	Forbidden                    = wrapFunc(403) // RFC 7231, 6.5.3
	NotFound                     = wrapFunc(404) // RFC 7231, 6.5.4
	MethodNotAllowed             = wrapFunc(405) // RFC 7231, 6.5.5
	NotAcceptable                = wrapFunc(406) // RFC 7231, 6.5.6
	ProxyAuthRequired            = wrapFunc(407) // RFC 7235, 3.2
	RequestTimeout               = wrapFunc(408) // RFC 7231, 6.5.7
	Conflict                     = wrapFunc(409) // RFC 7231, 6.5.8
	Gone                         = wrapFunc(410) // RFC 7231, 6.5.9
	LengthRequired               = wrapFunc(411) // RFC 7231, 6.5.10
	PreconditionFailed           = wrapFunc(412) // RFC 7232, 4.2
	RequestEntityTooLarge        = wrapFunc(413) // RFC 7231, 6.5.11
	RequestURITooLong            = wrapFunc(414) // RFC 7231, 6.5.12
	UnsupportedMediaType         = wrapFunc(415) // RFC 7231, 6.5.13
	RequestedRangeNotSatisfiable = wrapFunc(416) // RFC 7233, 4.4
	ExpectationFailed            = wrapFunc(417) // RFC 7231, 6.5.14
	Teapot                       = wrapFunc(418) // RFC 7168, 2.3.3
	MisdirectedRequest           = wrapFunc(421) // RFC 7540, 9.1.2
	UnprocessableEntity          = wrapFunc(422) // RFC 4918, 11.2
	Locked                       = wrapFunc(423) // RFC 4918, 11.3
	FailedDependency             = wrapFunc(424) // RFC 4918, 11.4
	TooEarly                     = wrapFunc(425) // RFC 8470, 5.2.
	UpgradeRequired              = wrapFunc(426) // RFC 7231, 6.5.15
	PreconditionRequired         = wrapFunc(428) // RFC 6585, 3
	TooManyRequests              = wrapFunc(429) // RFC 6585, 4
	RequestHeaderFieldsTooLarge  = wrapFunc(431) // RFC 6585, 5
	UnavailableForLegalReasons   = wrapFunc(451) // RFC 7725, 3

	InternalServerError           = wrapFunc(500) // RFC 7231, 6.6.1
	NotImplemented                = wrapFunc(501) // RFC 7231, 6.6.2
	BadGateway                    = wrapFunc(502) // RFC 7231, 6.6.3
	ServiceUnavailable            = wrapFunc(503) // RFC 7231, 6.6.4
	GatewayTimeout                = wrapFunc(504) // RFC 7231, 6.6.5
	HTTPVersionNotSupported       = wrapFunc(505) // RFC 7231, 6.6.6
	VariantAlsoNegotiates         = wrapFunc(506) // RFC 2295, 8.1
	InsufficientStorage           = wrapFunc(507) // RFC 4918, 11.5
	LoopDetected                  = wrapFunc(508) // RFC 5842, 7.2
	NotExtended                   = wrapFunc(510) // RFC 2774, 7
	NetworkAuthenticationRequired = wrapFunc(511) // RFC 6585, 6
)
