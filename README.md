<img width="50%" src="https://raw.githubusercontent.com/TboOffical/WitchServer/main/logo.png">

# WitchServer
The witch Web server, Written in golang<br>
I will only release the working, stable versions of of witch.<br>
If you want to try the unstable, you will have to build it.<br>
> Why is it named witch, welp https://randomwordgenerator.com/

# How to build witch
Windows / Mac / Linux
```bash
cd witch
go get -u github.com/gen2brain/dlgs
go build ./witch.go ./util.go ./wba.go ./listener.go
```
## New
# WBA
In the begining, I wanted witch to support PHP. But the I had a thought, what if witch
allowed people to write PHP like apps. But with Go instead of the... Intresting, Launage that people call PHP.
If you want to get started with It I recomend reading some of the Wiki entrys about it. 
**Example**
<br>
<img width="50%" src="https://i.imgur.com/74ePnNf.png">

# Witch DevServer
The easy way to get started with frontend web dev using the witch server
<img src="https://raw.githubusercontent.com/TboOffical/WitchServer/main/DevServer.png"><br>
Just add the path to your app, a port number and your set!

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
https://github.com/users/TboOffical/projects/1/views/1

## Cotainerisation

### Install && Usages

```bash
git clone https://github.com/mauricelambert/WitchContainer.git
docker build . -t witch:latest
docker run -d -p 8000:8000 witch:latest
```

#### Test container with DockerHub

```bash
docker run --rm -it -p 8000:8000 mauricelambert/witch:latest bash
```

#### Request using curl

```bash
curl 127.0.0.1:8000
```

### Link

 - https://github.com/mauricelambert/WitchContainer
