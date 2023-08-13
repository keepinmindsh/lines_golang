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

위와 같이 root 프로젝트에 설정후 shell을 통해서 아래의 명령어를 실행한다.

```
$ pre-commit install
```

이후 부터 commit 시점에 사전 체크해서 수정 및 배포가 가능하다.

### golangci-lint를 적용하고 싶다면,

아래의 링크에서 golangci-lint 설치 진행

- [https://golangci-lint.run/usage/install/](https://golangci-lint.run/usage/install/)

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
