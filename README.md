# simpsons-search

<img src="./img.png" width="300">

This web app allows users to search the transcripts of all Simpsons episodes from season 1 to 25 (included). The app is written in Go and uses a PostgreSQL database.

Getting the data (i.e. all the episode transcripts) was an interesting task in and of itself since I had to first find a website from which I could download all the transcripts. Fortunately I quickly found a site called transcripts.foreverdreaming.org that had all the transcripts. Even better, the transcript URLs contained sequential IDs, though sadly only up until the end of season 25. Regardless, this meant that I could easily write a Python program to download all the transcripts using a for loop. Once I had all the transcripts it was just a matter of cleaning up the HTML of each episode transcript, creating the database, and writing a simple Go app that takes advantage of Postgres's awesome full-text search functionality.

#### Link to project:

<a href="https://simpsons-search.herokuapp.com/" target="_blank">https://simpsons-search.herokuapp.com/</a>
