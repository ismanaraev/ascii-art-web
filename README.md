## ascii-art-web

### Description
This project hosts an http server that can generate ascii-art. You can chooose the fonts, available fonts are: standard, thinkertoy, shadow and doom. The server serves html template and there are two endpoints served:
1. GET / - returns main HTML page
2. POST /ascii-art returns formed ascii-art to the request. Request must have three fields: text - text of a banner to form, fonts - font to choose and width - width of a client's window, which is required to carry-over the ascii-art characters that overflow to the next line

### Authors
https://git.01.alem.school/ismanaraev

https://git.01.alem.school/aleka7sk

### Usage

run the server using:

` go run cmd/main.go `

if required, port can be changed either by editing config/config.json or setting environment variable "ASCII_WEB_PORT"

### Algorithm

The main logic of the ascii-art is located at pkg directory. There's a package that generates ascii-art independently from server. 
The server architechture is following clean-architechture rules. The main handlers and registry is located at internal/api/delievery.


