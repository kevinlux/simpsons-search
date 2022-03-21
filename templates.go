package main

const tplStrHome = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Simpsons Search</title>
    <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
  </head>
  <body style="background-image: url('/public/couch.png')">
    <div class="container">
	<div class="row" style="text-align: center">
		<h1 style="padding-top: 2em; font-family: monospace; font-size: 80px; background-color: rgba(255, 240, 0, 0.9);">Simpsons Search</h1>
		<h3 style="font-family: monospace; background-color: rgba(75, 119, 190, 0.9); color: white;">Find any quote from any episode between seasons 1 and 25 of the Simpsons</h3>
	</div>
	<div class="row" style="padding-top: 1.5em"><div class="col-sm-4 col-sm-offset-4">
		<form action="/" method="GET"><input class="form-control" autofocus name="q" maxlength=100 type="text"></form>
	</div></div>
		</div>
    </div>
  </body>
</html>
`

const tplStrResults = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Simpsons Search</title>
    <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
  </head>
  <body>
    <div class="container-fluid">
	<div class="row" style="text-align: center; background-color: #f1f1f1; padding: 1em 0">
<a href="/" class="col-sm-1">HOME</a>
		<div class="col-sm-4">
			<form action="/" method="GET"><input class="form-control" autofocus name="q" maxlength=100 type="text" value="{{.Query}}"></form>
		</div>
	</div>
	{{range .Results}}
	<div class="row" style="padding: 1.5em 0; border-bottom: 1px solid #eee"><div class="col-sm-11 col-sm-offset-1">
		<div style="font-size: 1.2em">
			<a href="/episode?e={{ .Title }}">{{ .Title}}</a>
		</div>
		{{.Snippet}}
	</div></div>
	{{end}}
    </div>
  </body>
</html>
`

const tplStrEpisode = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{.Title}}</title>
    <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
  </head>
  <body style="padding-left: 10px;">
		{{.Content}}
  </body>
</html>
`
