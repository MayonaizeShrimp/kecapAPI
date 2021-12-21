# kecapapi
TL;DR => API for kecap store made using Go Lang

What is this?
This is a simple web server created using gin-gonic for routing and gorm for easy database access.
To put it simply, it's a demo site to practice making websites. 

CONCERN:
This site, however, didn't implement middleware to check whether the user is authenticated.
The reason being that it didn't work for some reason(?). Several methods and libraries have been tried
to check user's session with no avail. 
Therefore, against the author's own judgment, the site's user security will be checked with browser cookies
using javascript.