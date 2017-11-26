{{ template "base.tpl" . }}

{{ define "content" }}
  <div class="row" style="margin-bottom:10px;">
    <div class="col-xs-12">
      <a href="{{urlfor "FoodController.NewFood"}}" role="button" class="btn btn-primary pull-right">Inserir nova</a>
    </div>
  </div>
  <div class="row">
    <div class="col-xs-12">
      <ul class="list-group">
        {{ range $_, $value := .Food_List }}
           <li class="list-group-item"> {{ $value.Name }} </li>
        {{ end }}
      </ul>
    </div>
  </div>

{{ end }}
