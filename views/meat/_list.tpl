{{ template "base.tpl" . }}
{{ define "content" }}
<div class="row">
    {{ range $_, $value := .Meal_List }}
    <div class="col-xs-4">
        <div class="panel panel-default">
            <div class="panel-heading">Refeição #{{$value.Id}}</div>
          <div class="panel-body">
                <p>{{ range $_, $food := $value.Foods }}
                    {{$food.Name}},
                {{end}}
                </p>
                <p>Data: {{date $value.Created "d/m/Y"}}</p>
                {{if $value.LactosePresence}}
                <p class="text-danger">Comeu algo com lactose.</p>
                {{end}}
          </div>
        </div>

    </div>
    {{end}}
</div>
{{ end }}