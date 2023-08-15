# Mutex 

## Index 

- [Read Lock](#read-lock)
- [Write Lock](#write-lock)

## Read Lock

읽기 Lock 사이에는 서로를 막지 않지만, 읽기 시도 중에 값이 바뀌면 안되므로 쓰기 Lock은 막음!

## Write Lock

쓰기 시도 중에 다른 곳에서 이전 값을 읽으면 안되고, 다른 곳에서 값을 바꾸면 안되므로 읽기, 쓰기 Lock 모두를 막음! 



# 참조 

> [https://fenderist.tistory.com/202](https://fenderist.tistory.com/202)   
> [https://gobyexample.com/mutexes](https://gobyexample.com/mutexes)  
