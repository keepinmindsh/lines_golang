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
      # - id: go-imports 위의 버전 repo에서 동작하지 않음 추후 버전 체크 필요
      - id: no-go-testing
      # - id: golangci-lint 위의 버전 repo에서 동작하지 않음 추후 버전 체크 필요
      #- id: go-unit-tests
```

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
