
# go-download #
Download a file package for Go
## Install ##
```
go get -u github.com/yisleyen/go-download
```
## Example ##
```go
err := download.DownloadFile("downloads", fileLink)
if err != nil {
	fmt.Println(err)
}
fmt.Println("Downloaded: " + fileLink)
```
