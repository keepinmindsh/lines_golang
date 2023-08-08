# Process, OS Thread, Goroutine  

- M : Machine 을 의미하며 OS Thread 이다. OS Thread 로 부르겠다.
- P : Logical Processor 를 의미하며 Context 로도 불린다. Process 로 부르겠다. 
- G : GoRoutine 이다.  

go routine 은 OS Thread 위에서 돌아가며 OS Thread 가 실행되려면 Process 하나를 잡고 있어야 한다.      

## 각 항목별의 설명 

- Processor 는 Logical CPU 를 의미하는데 보통 CPU 사양에서 4 core 8 thread 라고 할 때의 8 thread 가 의미하는 것이다. 
- OS Thread 는 OS가 스케줄링 해주는 하나의 프로그램 수행 공간이라 보면된다. 
- goroutine 은 Go runtime 이 스케쥴링 해주는 경량의 thread 라고 보면 된다. 

# Process

Context 라고도 한다. 코드가 실행되려면 CPU와 메모리 공간이 있어야 한다.   

- Process 개수는 runtime.GOMAXPROCS() 로 설정 가능하다.  
- logical CPU 개수는 runtime.NumCPU() 로 확인 가능하다. 
- runtime.GOMAXPROCS(runtime.NumCPU()) 라고 하면 
  - 프로세스의 개수를 컴퓨터가 실제로 병렬적으로 사용할 수 있는 최대치로 쓰겠다고 하는 것.

# OS Thread 

컴퓨터의 CPU Core 는 개수의 한계가 있다, 병렬적으로 돌릴 수 있는 프로그램의 개수 제한이 있다는 것이다.  

HW 적으로는 개수의 한계가 있지만,  
SW 적으로는 여러개의 Thread를 생성해서 프로그램 하다 할당해주는 것이다.  

# Goroutine 

고루틴은 OS Thread와 같지만, 다르다.  

- Go runtime이 관리해준다. 
  - 즉, OS Thread 는 OS가 생성하고 관리해주는데, 
  - Goroutine 은 Go runtime 이 생성하고 관리해주는 것이다. 

여기에서 n:m 이라는 표현이 나온다.      
1) 1:1 이라면 1개의 user level thread가 1개의 OS Thread 위에서 돌아가는 거다.   
2) n:1 이라면 n개의 user level thread가 1개의 OS Thread 위에서 돌아가는 거다.    
3) m:n 이라면 m개의 user level thread가 n개의 OS Thread 위에서 돌아가는 거다.    
→ 고언어는 goroutine 이라는 user level thread m개가 n 개의 OS Thread 위에서 돌아가는 것이다.     


# GOMAXPROCS 설정 

- GOMAXPROCS 환경설정  
  - GOMAXPROCS is a configuration parameter in Go that specifies the maximum number of OS threads that can execute Go code simultaneously.
- 기본값 세팅 
  - By default, GOMAXPROCS is set to the number of logical CPUs available on the machine, allowing Go to automatically utilize the available cores for parallel execution. 
- 성능 최적화 
  - Developers can adjust the value of GOMAXPROCS based on the specific workload and hardware characteristics to optimize the performance of their applications.
- Balancing Act
  - Setting GOMAXPROCS too high may lead to increased contention and context switching overhead. Finding the right balance is crucial for achieving optimal parallelism.

## P99 latency 

예를 들어 P99 45ms 이라면, Request를 100번 보냈을 때 99번은 45ms 보다 빨아야 한다(Should)는 뜻이다.    

특이 값이 얼마나 좋지 않은지 알아보려면 상위 백분위를 살펴보는 것도 좋다. 이때 사용하는 백분위는 95분위, 99분위, 99.9분위(p95, p99, p999). 요청의 p95, p99, p999가 특정 기준치보다 더 빠르면 해당 특정 기준치가 각 백분위의 응답 시간 기준치가 된다.
예를 들어 💡 95분위 응답 시간이 1.5초라면 100개의 요청 중 95개는 1.5초 미만이고, 100개의 요청 중 5개는 1.5초보다 더 걸린다.

- 예를 들어, 아마존은 💡 내부 서비스의 respnose time 요구사항을 p999로 기술한다. p999는 요청 1000개 중 1개만 영향이 있음에도 말이다. 보통 response time이 가장 느린 요청을 경험한 고객들은 많은 구매를 해서 고객 중에서 계정에 가장 많은 데이터를 갖고 있어서다. 반면 p9999를 최적화하는 작업에는 비용이 너무 많이 들어 아마존이 추구하는 목표에 충분히 이익을 가져다주지 못한다고 여겨진다. 최상위 백분위는 통제할수 없는 임의 이벤트에 쉽게 영향을 받기 때문에 response time을 줄이기가 매우 어려워 이점은 더욱 줄어든다.
- 예를 들어, 백분위는 💡 service level objective(서비스 수준 목표, SLO)와 service level agreement(서비스 수준 협약서, SLA)에 자주 사용하고, 기대 성능과 서비스 가용성을 정의하는 계약서에도 자주 등장한다.

# 참고 링크 

> [https://jusths.tistory.com/142](https://jusths.tistory.com/142)   
> [Memory FootPrint / 메모리 사용량](https://en.wikipedia.org/wiki/Memory_footprint)    
> [Performance Bottlenecks go apps](https://engineering.grab.com/performance-bottlenecks-go-apps)     
> [runtime: make the proportion of CPU the GC uses based on actual available CPU time and not GOMAXPROCS](https://github.com/golang/go/issues/59715)      
> [Surprise at CPU Hogging in Golang](https://winder.ai/cpu-hogging-in-golang/)   
> [Optimizing a Golang Service to Reduce Over 40% CPU](https://coralogix.com/blog/optimizing-a-golang-service-to-reduce-over-40-cpu/)