# WitchServer
The witch Web server, Written in golang<br>
This is a work-in-progress. So code might be a bit messy.<br>
I will only release the working, stable versions of of witch.<br>
If you want to try the unstable, you will have to build it.<br>
> Why is it named witch, welp https://randomwordgenerator.com/

# How to build witch
Windows / Mac / Linux
```bash
cd witch
go get -u github.com/gen2brain/dlgs
go build ./witch.go ./util.go ./listener.go
```
# Witch DevServer
The easy way to get started with frontend web dev using the witch server
<img src="https://raw.githubusercontent.com/TboOffical/WitchServer/main/DevServer.png">

# Information
The server will look for a index.html
file when "/" is acessed so make sure to have that created<br>
Nothing happens if you dont, you just get a error

You can create a witch cofig file by createing the file<br>
witch.json in the dirictory that the server exe is in.

```json
{
  "/route" : "file.html",
  "/whatever" : "sdfsdf.html"
}
```
More options comeing in the future for things like
POST requests<br>

You can also add a ssl certificate now by creating a file
named cert.json<br>
Once created, put in the location of your cert and key file and witch will do the rest
```json
{
    "enableTLS": true,
    "crt_file": "localhost.cert",
    "key_file": "localhost.key"
}
```
if witch dose not find the file, it will start with no ssl

# Benchmarks

Tested on template : https://templatemo.com/tm-565-onix-digital<br>
Its a big templates with lots of content to load in

<table>
  <thead>
    <tr>
      <th>Server</th>
      <th>Speed</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>Witch</td>
      <td>2.19s</td>
    </tr>
    <tr>
      <td>Nginx</td>
      <td>3.11</td>
    </tr>
  </tbody>
</table>

# ToDo List

- [x] Basic File loading and routeing
- [x] custom routes based on json file
- [ ] php? or altertive
- [x] HTTPS!
- [ ] Status gui
- [ ] templateing based on json

This list is going to grow, and fast. <br>
So if there is a missing feature it will be on here soon
