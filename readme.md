# Memory Leak web app example

## An example of a memory leak in a golang application


### How to run:

```shell
docker run -p 8080:8080 --rm --memory=500m --memory-swappiness=0  nathanmartins/leaky
```