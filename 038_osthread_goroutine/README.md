# Process, OS Thread, Goroutine  

- M : Machine을 의미하며 OS Thread 이다. OS Thread로 부르겠다.
- P : Logical Processor 를 의미하며 Context 로도 불린다. Process로 부르겠다. 
- G : GoRoutine 이다.  

go routine은 OS Thread 위에서 돌아가며 OS Thread가 실행되려면 Process 하나를 잡고 있어야 한다.      

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


# GOMAXPROCS Setting 

- GOMAXPROCS Configuration 
  - GOMAXPROCS is a configuration parameter in Go that specifies the maximum number of OS threads that can execute Go code simultaneously.
- Default Setting
  - By default, GOMAXPROCS is set to the number of logical CPUs available on the machine, allowing Go to automatically utilize the available cores for parallel execution. 
- Performance Optimization
  - Developers can adjust the value of GOMAXPROCS based on the specific workload and hardware characteristics to optimize the performance of their applications.
- Balancing Act:
  - Setting GOMAXPROCS too high may lead to increased contention and context switching overhead. Finding the right balance is crucial for achieving optimal parallelism.

# 참고 링크 

> [https://jusths.tistory.com/142](https://jusths.tistory.com/142)   
> [Memory FootPrint / 메모리 사용량](https://en.wikipedia.org/wiki/Memory_footprint)   