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
