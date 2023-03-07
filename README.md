<p align="center"><img src="https://user-images.githubusercontent.com/1092882/60883564-20142380-a268-11e9-988a-d98fb639adc6.png" alt="webgo gopher" width="128px"/></p>

# Youtube-Data-Golang

This service contains API to fetch latest video sorted in reverse chronological order of their publishing date-time from YouTube for a given tag/search query in a paginated response.

### Functionalities
  - It runs background job to fetching latest videos data from Youtube and pushes into database in every X configured time.
  - Youtube searches response return in sorted reverse chronological order of their publishing date-time.
  - It has exposes APIs to get Youtube data from database in a paginated response sorted in descending order of published datetime.
  - A search API to search the stored videos using their title and description.
  
### Dependencies to run this project
  - Need to set `credentials.json` file path. Please create a new project on [google-cloud-platform](https://developers.google.com) and download the credential file.
  - Need to set mongodb config for database operations.
  

### Problem breakdown (High Level)

I have devided this problem statements into some steps:
  - Getting data from Youtube
    - First step, getting search data from Youtube.
    - I have implemented the Youtube APIs which calls the Youtube services to perform the search and return the response.
    - Implemented response sorting in reverse chronological order of their publishing date-time.
    - Reference
      - [youtube-golang-pkg](https://pkg.go.dev/google.golang.org/api/youtube/v3)
      - [youtube-api-doc](https://developers.google.com/youtube/v3/getting-started)
  - CRON job to getting data from Youtube and push data into database.
    - Second step, for calling the Youtube API continuously in background with some X interval, implemented the CRON Job.
    - Reference
      - [cron-job-pkg](https://github.com/robfig/cron)
  - Database operations for adding the youtubes search results
    - Third step, After getting data from youtube need to insert into database.
    - Implemented mongodb connections and operations to insert bulk data into database.
    - Reference
      - I have already built mongodb wrapper library. [go-mongo-lib](https://github.com/lokesh-go/go-mongo-lib)
  - Now time to impelement Restful APIs for performing GetData and SearchData from database.
    - Implemented an the server to handle the restful APIs.
    - Reference
      - [echo-framework](https://echo.labstack.com/)

### How can scale this project for higher traffic
  - For database
    - Have already added connection pooling, timeouts etc. just need to tunning these values according to the traffic.
    - Bulk insertion operation will be heavy operation if inserting the lots of the bulk data.
      - It can be optimised when data will be insert in batch by batch.
    - If searches and Get APIs received too much traffic then will have to consider the Caching for better API response time and preventing from too much load to database.
  - For Youtube APIs
    - As hitting the Youtube APIs within X time duration so chances to quota limit will be exhausted, So either can increase the X time interval value or need to add the multiple account handing.

### Project best practices
- Implemented clean golang project directory architecture.
- Added comments and meaningful variable/function names which make project easy understandable.
- Added dependency injection.
- All pkgs kept separated.
- Error logging at the parent level.
- Designed extendable project so will add more functionality easily.