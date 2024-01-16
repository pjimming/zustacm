api:
	goctl api go -api core/core.api -dir core --home template --style go_zero
	goctl api format --dir core/apis -declare

gentool:
	gentool -c gen/gen.yaml

dao:
	@go run gen/cmd/gen.go -model=$(model) -target=core/dao -home=gen/dao.tpl