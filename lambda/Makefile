.PHONY: build
build: draas.go
	@echo "==> Building DRAAS..."
	GOARCH=amd64 GOOS=linux go build -v -ldflags="-s -w" $^

.PHONY: clean
clean:
	@rm draas draas.zip

.PHONY: zip
zip: draas
	@echo "==> Creating zipfile..."
	@zip draas.zip draas

# IAM policy functions
.PHONY: iam-create
iam-create: lambda-trust-policy.json
	aws --endpoint-url=http://localhost:4593 iam create-role \
	--role-name lambda-basic-execution \
	--assume-role-policy-document file://$^
.PHONY: iam-attach
iam-attach:
	aws --endpoint-url=http://localhost:4593 iam attach-role-policy \
	--role-name lambda-basic-execution \
	--policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
.PHONY: iam-get
iam-get:
	aws --endpoint-url=http://localhost:4593 iam get-role \
	--role-name lambda-basic-execution

# Lambda deployment jobs
.PHONY: delete
delete:
	-aws --endpoint-url=http://localhost:4574 lambda delete-function \
	--function-name draas_go
.PHONY: deploy
deploy: draas.zip
	awslocal lambda create-function \
	--function-name draas_go \
	--zip-file fileb://$^ \
	--handler draas \
	--runtime go1.x \
	--role "arn:aws:iam::000000000000:role/lambda-basic-execution"
.PHONY: redeploy
redeploy: delete deploy
	@echo "--> Redeployed zipfile."
	@true

.PHONY: invoke
invoke:
	@aws --endpoint-url=http://localhost:4574 lambda invoke \
	--function-name draas_go \
	--invocation-type "RequestResponse" \
	--payload '{"size": 20, "count": 4}' \
	/dev/tty

.PHONY: localstack
localstack:
	@echo "==> Starting Localstack..."
	@localstack start
