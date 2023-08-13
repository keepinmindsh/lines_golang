# Pre Commit 설정하기

Install PreCommit

```shell
brew install pre-commit
```

Make file ( .pre-commit-config.yaml ) on your root project

```yaml
repos:
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.3.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-yaml
      - id: check-added-large-files
  - repo: https://github.com/dnephin/pre-commit-golang
    rev: v0.5.0
    hooks:
      - id: go-fmt
      - id: go-imports
      - id: no-go-testing
      # - id: golangci-lint 위의 버전 repo에서 동작하지 않음 추후 버전 체크 필요
      #- id: go-unit-tests
```

위와 같이 root 프로젝트에 설정후 shell을 통해서 아래의 명령어를 실행한다.

```
$ pre-commit install
```

이후 부터 commit 시점에 사전 체크해서 수정 및 배포가 가능하다.

### golangci-lint를 적용하고 싶다면,

아래의 링크에서 golangci-lint 설치 진행

- [https://golangci-lint.run/usage/install/](https://golangci-lint.run/usage/install/)

### go-imports를 적용하고 싶다면,

아래의 가이드에 따라서 go-import 설치 후 사용할 것

- [https://pkg.go.dev/golang.org/x/tools/cmd/goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)

> 참고로 multimodule로 구성된 repo의 경ㅎ우 root project에서 현재 동작하지 않는 케이스가 존재하므로 참고 할 것

아래의 케이스 참고할 것

```
ERROR [linters_context] typechecking error: pattern ./...: directory prefix . does not contain modules listed in go.work or their selected dependencies
```

### 그외 기타 규칙 적용하고 싶다면 github 링크에서 확인후 수정 바람.

- [https://github.com/dnephin/pre-commit-golang](https://github.com/dnephin/pre-commit-golang)

## 만약, .pre-commit-config.yaml 를 추가로 수정하는 경우

```
git commit -m "osthread - gorouting"
[ERROR] Your pre-commit configuration is unstaged.
`git add .pre-commit-config.yaml` to fix this.
```

위의 설명 처럼 git add 를 통해서 아래와 같이 추가하여 커밋해야한다.

```shell
$ git add .pre-commit-config.yaml
```
