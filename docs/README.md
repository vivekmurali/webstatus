# WEBSTATUS



### Overview:

User creates an account and adds a list of APIs or webpages that the service sends a request to every x amount of time. If the response is a 200 response, then the logs are collected and displayed when requested. 

If the response is a 400 or a 500 response, an email notification is sent with a link to the logs. 



### Process:

A worker process runs the cycle every 5 minutes, if the current time is greater than the interval + last_check then a separate goroutine is spawned and get request performed for the same. 

` CURRENT_TIME - LAST_CHECK > INTERVAL `

Make a get request to the URL, if the URL responds with a SUCCESS response, write response and stuff to the logs, if the URL responds with a FAILURE, send an email with the response to the user’s email ID.

When a new record is inserted, that record starts from the next cycle.



### APIs:

##### Public:

* Register
  * Creates a user
* Login
  * Logs in and returns a token
  * Log in keeps a track of session in a database
* AddURL
  * Adds a URL to the database under user's account
* RemoveURL
  * Remove URL from the database and add it to the removed database where it will be removed automatically after a week
* GetSingleData
  * Get data for a single website
* GetAllData
  * Get just the status codes for all websites

##### Private:

* Networking
  * Data usage and memory info



### Database:

* User
  * ID
  * Username
  * Password
  * Email ID
  * Account creation time
  * Account level
* Web
  * {user.ID}
  * ID
  * Creation time
  * Request type
  * URL
  * BODY
  * Interval
* Logs
  * {web.ID}
  * ID
  * Time
  * Logs
* Deleted User
* Deleted Web



### Ideas:

* Limit user to 5 URLs for free account.
* Have users choose URL's that can be public and will have their own links, so anyone can see status
* Instant check of status on homepage for non registered users
* Status codes help cheatsheet
* Have a graph with status codes of last 10 checks

### Code references:

##### Time query:

```sql
SELECT * FROM NEWTEST WHERE CREATE_TIME NOT BETWEEN LOCALTIME - INTERVAL '5 MINUTES' AND LOCALTIME;
```

##### Time query with interval:

```sql
SELECT * FROM NEWTABLE WHERE CREATE_TIME NOT BETWEEN LOCALTIME - CAST(INTER AS INTERVAL) AND LOCALTIME;
```

