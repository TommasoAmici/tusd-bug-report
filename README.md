# Usage

Run with

```bash
go run ./cmd/main.go
```

# Issue

**Describe the bug**
I'm using `tusd` programmatically in my go backend behind nginx, and I keep getting this error:

```
Error: tus: invalid or missing offset value, originated from request
(method: HEAD, url: https://mydomain.tld/files/b53af271868fdc811f4f3a93d6f5a7c9,
response code: 200, response text: , request id: n/a)
```

I can upload small files (<10mb), but when trying bigger files only the first chunk is saved and then the error occurs.

When looking at the headers, it seems like the tus headers are not there.

e.g. on a HEAD response I get this. I would expect to see `Upload-Offset` and `Upload-Length` in there.

```http
HTTP/2 200 OK
server: nginx/1.19.6
date: Sun, 17 Jan 2021 14:36:27 GMT
content-type: application/zip
content-length: 644676112
content-disposition: attachment;filename="filename.zip"
tus-resumable: 1.0.0
x-content-type-options: nosniff
x-cache-status: MISS
x-frame-options: SAMEORIGIN
x-xss-protection: 1; mode=block
referrer-policy: no-referrer-when-downgrade
strict-transport-security: max-age=31536000; includeSubDomains
X-Firefox-Spdy: h2
```

**To Reproduce**
I've made a repository with a (not) working example. See here: https://github.com/TommasoAmici/tusd-bug-report

**Expected behavior**
I should be able to upload files on the endpoint.

**Setup details**
Please provide following details, if applicable to your situation:
- Operating System: macOS for local development, Ubuntu server for staging
- Used tusd version: `1.4.0`
- Used tusd data storage:
```go
        store := filestore.FileStore{
		Path: *app.tusDir,
	}

	composer := tusd.NewStoreComposer()
	store.UseIn(composer)
```
- Used tusd configuration: 
```go
       handler, err := tusd.NewHandler(tusd.Config{
		BasePath:                "/files/",
		StoreComposer:           composer,
		MaxSize:                 8 * 1_000_000 * 1_000,
		RespectForwardedHeaders: true,
		NotifyCompleteUploads:   true,
		Logger:                  app.infoLog,
		PreUploadCreateCallback: app.checkIfOrderExistsTusHook,
	})
```
Then I integrate it in my mux like this:
```go
mux.Handle("/files/", http.StripPrefix("/files/", handler))
```
- Used tus client library: Uppy.js
```ts
new Uppy<StrictTypes>({
    id: "client-upload",
    debug: true,
    autoProceed: false,
    restrictions: {
      minFileSize: 1,
      maxFileSize: 8 * GB,
      maxNumberOfFiles: numFiles,
      minNumberOfFiles: numFiles,
      allowedFileTypes: uppyFileTypes,
    },
    locale: languageMap[language],
    meta: {
      orderID: orderID,
      orderUUID: orderUUID,
      orderEmail: orderEmail,
      admin: false,
    },
  })
```
