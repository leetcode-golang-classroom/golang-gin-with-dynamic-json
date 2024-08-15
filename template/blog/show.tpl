{{ define "blog/show.tpl" }}
    {{ template "layout/header.tpl" .}}
        <h1 class="card-title">{{ .blog.Title }}</h1>
        <p class="card-text">{{ .blog.Content }}</p>
    {{ template "layout/footer.tpl" .}}
{{ end }}