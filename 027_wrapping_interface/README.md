# https://go.dev/blog/go-cloud

MultiCloud 를 동적으로 구현하여 호출할 수 있는 프로세스 구조 정의 

### 기본 구조 검토 

- Strategy Pattern 및 Abstract Factory Pattern을 적용 
  - .env 파일에 설정된 정보를 기준으로 gcp, aws, azure 등을 동적으로 객체를 생성하는 구조로 적용 

- Setting 방식 
  - .env 를 제어할 수 있는 간단한 GUI 를 통해서 env 설정 변경 
    - github 에 최종 소스 관리 ( .env 파일 관리 )
    - github 로 부터 .env 파일을 읽어와서 해당 gui 프로그램을 통해서 설정 변경할 수 있도록 처리 
      - LOCAL 에서 imports 가 가능하도록 처리해도 무방함.

### Electron App 만들기 

- [Electron App 만들기](https://blog.codefactory.ai/electron/create-desktop-app-with-react-and-electron/1-project-setting/) 
- [Electron Gui 만들기](https://github.com/electron-react-boilerplate/electron-react-boilerplate)
- https://www.youtube.com/watch?v=zq-XcnjLpXI
- https://www.youtube.com/watch?v=BbZmLXBDGnU


### Golang Framework 

- https://www.youtube.com/watch?v=10miByMOGfY 