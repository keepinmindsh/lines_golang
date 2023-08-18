# GC in Golang

## Go's Garbage Collection

- Concurrent Mark and Sweep GC
- Use hybrid write barrier for concurrency
- non-generational & node-compacting GC
- Pager & soft memory limit for trigger GC
- Non-Cooperative Preemption for Latency

### Concurrent Mark And Sweep (CMS)

동시성 환경에서 사용하는 Mark & Sweep 에 대한 처리

### Tri-color abstraction

ref : [On-the-Fly Garbage Collection: An Exercise in Cooperation](https://lamport.azurewebsites.net/pubs/garbage.pdf)


**black**은 접근할 수 있는 객체를 의미하며 지워져서는 안됩니다.
**grey**는 black 과 유사하지만 다른데, root로 부터 접근은 가능하지만 아직 모든 자식 객체들이 mark되지 않아서 mark 단계가 끝나지 않았음을 암시합니다.
**white** 는 mark 되지 않는 객체를 의미합니다. sweep 단계에서 단계에서 white 객체들을 해제하게 됩니다.

tri-color abstraction은 다음의 불변식을 만족합니다.
- mark가 끝났으면 black 객체로 부터 white 객체들을 전부 접근이 불가능한 객체입니다.
- mark가 끝났으면 grey 객체가 없습니다.

### Atomic Operation

원자성 처리로 인한 처리를 위해서 Stop The World를 실행시키는 구조가 됨 - Mark & Sweep

STW는 GC를 수행하기 위해서 프로그램을 정지시키는 것을 의미합니다.
STW는 시스템의 SLO 지표에 영향을 줄 수 있으며, 때때로 치명적일 수 있습니다.

> [SLO, SLI](https://badcandy.github.io/2018/12/28/SRE-chapter04/)

### GC의 Latency를 고려하기 위해서

- [CMS](Concurrent Mark & Sweep)

CMS 를 Golang 에서는 고려하고 잇는데, STW 는 Latency 에는 좋지 않고,
병렬적으로 수행할 수 있는 방법이 있다면 GC를 스레드 또는 고루틴 내에서 동작하게 하는 것!

### Correctness of Concurrent Collector

- Safefy : Property might be that the garbage collector will never collect an object that is still in use
- Liveness : Property might be that it will eventually collect all garbage

### Write Barrier 

Write Barrier - 특정 객체에 대한 참조가 생기는 경우에 수정을 가하는 것 

### Memory Allocation 

- non-coping 
- non-moving 

#### Memory Fragmentation 

메모리를 할당하고 헤제하는 과정에서 조각조각 난 메모리들이 잘 관리되지 않아 필요한 메모리보다 더 많이 사용하게 되는 현상 

> [용어 - 메모리 단편화](https://m.blog.naver.com/PostView.naver?isHttpsRedirect=true&blogId=ljc8808&logNo=220303236020)

#### Read Barrier for moving GC 

#### TCMalloc 

- 스레드 캐쉬를 사용하는 메모리 할당자를 이용해서 메모리 단편화를 해결하고자 함. 

### P, M, G 

#### Goroutine의 스택 

### Memory Allocation 

- 작은 객체 : 각각 메모리사이즈별로 각각의 사이즈를 지정하게 되고 지정된 메모리끼리 각자 할당되게 되어 있음. 
메모리 단편화 문제를 해소할 수 잇는 부분이 있음. 
- 큰 객체 : 힙에서 관리함. 

### Mark Bit, FreeList, Bitmap 

- Use Bitmap 
- Mark Bit
- FreeList 

### Triggering GC 

- Tradeoff GC
  - Trade Off 가 있을 수 있음 

GC를 적절하게 증가시키는 방법에 대해서 고려해야함. GC 주기에 따라서 CPU 사용률에 영향을 미치는 데, 아래의 항목에 제어 파라미터이다.  

- GOGC 
- Soft Memory Limit 

> Help GC for memory allocation.      
> force gc period ( 2min ).    

### Performance 

- Spawn per-P background worker with GOMAXPROCS 
- Target 25% CPU Resources 
- User Goroutine assist GC mark works 
- CPU Limiter 

### Preemption for STW 

- Scheduler Call Preemption


> [2023 고퍼콘 2일차](https://www.youtube.com/watch?v=8AUVKh0qJgU&t=23357s)