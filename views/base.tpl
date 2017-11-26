<!DOCTYPE html>
<html>
<head>
    <title>Lin Li</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap.min.css">
    <link rel="stylesheet" href="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap-theme.min.css">
     {{ block "css" . }}{{ end }}
</head>
<body>
    <nav class="navbar navbar-default">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="{{urlfor "MealController.List"}}">Controle de Refeição</a>

        </div>

        <ul class="nav navbar-nav navbar-right" style="margin-right:10%">
        <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Funcionalidades<span class="caret"></span></a>
          <ul class="dropdown-menu">
            <li><a href="{{urlfor "MealController.New"}}">Registrar refeição</a></li>
            <li><a href="{{urlfor "FoodController.New"}}">Registrar alimento</a></li>
          </ul>
        </li>
      </ul>
    </nav>
    <div class="container">
        {{ block "content" . }}{{ end }}
    </div>
    <script type="text/javascript" src="http://code.jquery.com/jquery-2.0.3.min.js"></script>
    <script src="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/js/bootstrap.min.js"></script>
     {{ block "js" . }}{{ end }}
</body>
</html>