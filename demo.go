package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	// Once changes are merged in, this should be "code.google.com/p/goauth2/oauth"
	"github.com/CrossRef/goauth2-orcid/oauth"
)

var (
	clientId       = "XXXX-XXXX-XXXX-XXXX"                  // FILL ME IN
	clientSecret   = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX" // FILL ME IN
	scope          = "/orcid-profile/read-limited"
	redirectUrl    = "https://prospect.labs.crossref.org/auth/orcid/callback"
	prodAuthUrl    = "https://orcid.org/oauth/authorize"
	sandboxAuthUrl = "http://sandbox-1.orcid.org/oauth/authorize"
	tokenUrl       = "https://api.orcid.org/oauth/token"

	// The ORCID id is inserted once it comes back from auth.
	requestUrl = "https://api.orcid.org/%s/orcid-profile"
	code       = flag.String("code", "", "Authorization Code")
	cachefile  = "cache.json"
)

func main() {
	flag.Parse()

	var authUrl = prodAuthUrl

	// Set up a configuration.
	config := &oauth.Config{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  redirectUrl,
		Scope:        scope,
		AuthURL:      authUrl,
		TokenURL:     tokenUrl,
		TokenCache:   oauth.CacheFile(cachefile),
	}

	// Set up a Transport using the config.
	transport := &oauth.Transport{Config: config}

	// Try to pull the token from the cache; if this fails, we need to get one.
	token, err := config.TokenCache.Token()
	if err != nil {
		if clientId == "" || clientSecret == "" {
			flag.Usage()
			fmt.Fprint(os.Stderr, err) // Changed
			os.Exit(2)
		}
		if *code == "" {
			// Get an authorization code from the data provider.
			// ("Please ask the user if I can access this resource.")
			url := config.AuthCodeURL("")
			fmt.Println("Visit this URL to get a code, then run again with -code=YOUR_CODE\n")
			fmt.Println(url)
			return
		}
		// Exchange the authorization code for an access token.
		// ("Here's the code you gave the user, now give me a token!")
		token, err = transport.Exchange(*code)
		if err != nil {
			log.Fatal("Exchange:", err)
		}
	}

	// Make the actual request using the cached token to authenticate.
	// ("Here's the token, let me in!")
	transport.Token = token

	getInfoUrl := fmt.Sprintf(requestUrl, token.Extra["orcid"])
	fmt.Println("Fetch", getInfoUrl)

	// Make the request.
	r, err := transport.Client().Get(getInfoUrl)
	if err != nil {
		log.Fatal("Get:", err)
	}
	defer r.Body.Close()

	// Write the response to standard output.
	io.Copy(os.Stdout, r.Body)

	// Send final carriage return, just to be neat.
	fmt.Println()
}
