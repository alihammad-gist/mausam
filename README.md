# Mausam

CLI wrapper around wttr.io weather service. wittr.io has inaccurate
location resolver, so this app uses ipinfo.io service. Add ipinfo.io
token before usage. ipinfo allow 50000 requests per month for free. The
token is available immediately after signup, assign your token in the main.go
file's IpinfoToken variable.
