# Go ORCID

Example code for interfacing with [ORCID](http://orcid.org) in Go using the `goauth2` library. The library does nearly all of the work so this is a demo of how to authenticate against ORCID rather than an extra layer of library. 

This will log in and fetch the ORCID id for a user and retrieve their basic bibliographic details.

This file is heavily based on the demo file that comes with `goauth2`. It currently requires a forked version of `goauth2` to get the ORCID ID out of the response but hopefully the changes will be merged to the main library soon.

## How to get started:

1. Install dependencies

       go get github.com/CrossRef/goauth2-orcid/oauth

   In future, when the changes are merged back, this will be:


       go get code.google.com/p/goauth2/oauth

2. Insert your credentials:

   - `clientId` your OAuth client ID
   - `clientSecret` and secret
   - `scope` the scope you wish to use
   - `redirectUrl` the URL you want OAuth to redirect the user to
   
   All of these should be agreed between you and ORCID. The ORCID API docs are comprehensive.
   
   Both the sandbox and production URLs are included. Swap out `authUrl`.
   
   As it stands, the code will fetch the ORCID id and then request the basic profile information. Change the value of `requestUrl` if you want to do [something different](http://support.orcid.org/knowledgebase/articles/120162-orcid-scopes).

   Normal OAuth2 development rules apply. You can put the domain of your `redirectUrl` in your own hosts file.

3. Run it

 
    1. First run:

            go run demo.go

      and follow instructions. 
    
    2. First run will return a URL to authenticate which you should paste into your browser.
    3. On success you will will end up with an ORCID redirect including a code.
    4. Run again with the code:
    
            go run demo.go --code=MYCODE
       
       and you should see the ORCID details returned to you.
    
    If you're playing around with the API, remember to delete `cache.json` each time otherwise the code won't try and re-authenticate. This is `.gitignored`.
