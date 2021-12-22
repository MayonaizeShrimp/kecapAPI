# Kecap API
TL;DR => API for kecap store made using Go Lang

**What is this?**
This is a simple web server created using gin-gonic for routing and gorm for easy database access. It was also using github.com/gin-contrib/cors to handle CORS requests, but it didn't work for some reason.
 
To put it simply, it's a demo site to practice making websites.


**CONCERNS:**
This site, however, didn't implement middleware to check whether the user is authenticated.
The reason being that it didn't work for some reason(?). Several methods and libraries have been tried
to check user's session with no avail. 
Therefore, against the author's better judgment, there will be no security on this API