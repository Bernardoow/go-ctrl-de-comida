{{ template "base.tpl" . }} {{ define "content" }}
<div class="row">
    <div class="col-xs-12">
        <form method="POST">
            <div class="form-group {{if index .errors "name"}}has-error{{end}}">
                <label for="foodName">Nome:</label>
                <input type="text" class="form-control" id="Name" name="name" placeholder="Digite o nome da comida">
                {{if index .errors "name"}}<label class="control-label" for="inputError1">{{index .errors "name"}}</label>{{end}}
            </div>
            <button type="submit" class="btn btn-primary">Gravar</button>
        </form>
    </div>
</div>
{{ end }}