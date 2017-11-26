{{ template "base.tpl" . }} {{ define "content" }}
<div class="row">
    <div class="col-xs-12">
        <form method="POST">
            <fieldset>
                <legend>Alimentos:</legend>
                {{ range $_, $value := .Food_List }}
                <div class="form-group col-xs-4">
                    <label for="cb{{$value.Id}}">
                        <input id="cb{{$value.Id}}" value="{{$value.Id}}" name="Foods" type="checkbox"> {{$value.Name}}
                    </label>
                </div>
                {{end}}
                <hr>
                <div class="row">
                    <div class="col-xs-12">
                        <div class="checkbox">
                            <label>
                                <input name="LactosePresence" type="checkbox"> Algum alimento possui lactose?
                            </label>
                        </div>
                    </div>
                </div>
            </fieldset>
            <button type="submit" class="btn btn-primary">Gravar</button>
        </form>
    </div>
</div>
{{ end }}