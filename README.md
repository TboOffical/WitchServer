# WitchServer
The witch Web server, Written in golang<br>
This is a work-in-progress. So code might be a bit messy.<br>
I will only release the working, stable versions of of witch.<br>
If you want to try the unstable, you will have to build it.<br>
> Why is it named witch, welp https://randomwordgenerator.com/

# How to build witch
Windows / Mac / Linux
```
cd witch
go get -u github.com/gen2brain/dlgs
go build ./witch.go
```

# Information
The server will look for a index.html
file when "/" is acessed so make sure to have that created<br>
Nothing happens if you dont, you just get a error

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
- [ ] custom routes based on json file
- [ ] php? or altertive
- [ ] HTTPS!
- [ ] Status gui
- [ ] templateing based on json

This list is going to grow, and fast. <br>
So if there is a missing feature it will be on here soon
