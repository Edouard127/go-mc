package microsoft

const (
	LiveRedirectURI  = "https://login.live.com/oauth20_desktop.srf"
	LiveTokenRefresh = "https://login.live.com/oauth20_token.srf"
)

const (
	XboxLiveAuthorizationEndpoint = "https://user.auth.xboxlive.com/user/authenticate"
	XboxXSTSAuthorizationEndpoint = "https://xsts.auth.xboxlive.com/xsts/authorize"
	XboxLiveAuthHost              = "user.auth.xboxlive.com"
	XboxLiveAuthRelay             = "http://auth.xboxlive.com"
	XboxLiveClientID              = "000000004C12AE6F"
	XboxLiveScope                 = "service::user.auth.xboxlive.com::MBI_SSL"
)

const (
	MicrosoftOauth2Host            = "https://login.microsoftonline.com/consumers/oauth2/v2.0"
	MicrosoftAuthorizationEndpoint = MicrosoftOauth2Host + "/authorize"
	MicrosoftDeviceCodeEndpoint    = MicrosoftOauth2Host + "/devicecode"
	MicrosoftTokenEndpoint         = MicrosoftOauth2Host + "/token"
	MicrosoftNativeClient          = MicrosoftOauth2Host + "/nativeclient"
	MicrosoftScope                 = "XboxLive.signin offline_access"
	MicrosoftClientID              = "88650e7e-efee-4857-b9a9-cf580a00ef43"
)

const (
	MinecraftAuthorizationEndpoint = "https://api.minecraftservices.com/authentication/login_with_xbox"
	MinecraftProfileEndpoint       = "https://api.minecraftservices.com/minecraft/profile"
	MinecraftCertificateEndpoint   = "https://api.minecraftservices.com/player/certificates"
	MinecraftAuthRelay             = "rp://api.minecraftservices.com/"
)

const (
	DefaultCacheFile   = "auth.cache"
	DefaultKeyPairFile = "keypair.pem"
)
