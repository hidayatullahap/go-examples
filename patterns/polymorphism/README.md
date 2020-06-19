Polymorphism is an object-oriented programming concept that refers to the ability of a variable, function or object to take on multiple forms. In a programming language exhibiting polymorphism, class objects belonging to the same hierarchical tree (inherited from a common parent class) may have functions with the same name, but with different behaviors. [[1]](https://www.educative.io/edpresso/what-is-polymorphism)

For this example we will create API, DB, and Mock services with same interface. We can call all same function with those 3 implementation.
 
 ![alt text](https://i.imgur.com/h5Bk2Q9.jpg "example")
 
 In this illustration we can switch between multiple service implementation. For example we need to test API hit without
 actually hit the API, we can use mock instead. Code available in package service and service_example_test.go   
