# Deeptrace


A web service implementing FizzBuzz with a Pink Flamingo and Roman Calculator 

# How To Run Locally

docker build -t deeptrace .

docker run -p 8081:8081 deeptrace

# How To Test
 

[1] FizzBuzz with a Pink Flamingo
```
curl -X GET 'http://localhost:8081/pinkflamingo?from=0&to=10'

Expected Response: 
["Pink Flamingo","Flamingo","Flamingo","Flamingo","4","Flamingo","Fizz","7","Flamingo","Fizz"]
```

[2] Roman Calculator
```
curl -X POST 'http://localhost:8081/romancalc' -d "{  \"expr\": \"II * VIII / II\"}"

Expected Response: VIII
```


# API Documentation
a) Open http://localhost:8081/api/ in the web browser. You can also test there. 

