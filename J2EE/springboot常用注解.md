### springboot常用注解



-  **@SpringBootApplication：**

```java
申明让spring boot自动给程序进行必要的配置，这个配置等同于：
@Configuration  @EnableAutoConfiguration @ComponentScan 三个配置。 
```

- **@ResponseBody**

```java
直接返回json格式:
表示该方法的返回结果直接写入HTTP response body中，一般在异步获取数据时使用，用于构建RESTful的api。在使用@RequestMapping后，返回值通常解析为跳转路径，加上@responsebody后返回结果不会被解析为跳转路径，而是直接写入HTTP response body中。比如异步获取json数据，加上@responsebody后，会直接返回json数据。该注解一般会配合@RequestMapping一起使用
```
- **@RequestBody **
```java
@RequestBody 传递的参数变为json
```

- **@Value**

```java
注入application.properties/application.yml配置的属性的值。

@Value(value = “#{message}”) 
private String message;
```

- **@PathVariable**

```java
路径参数 @RequestMapping(/{id}/{pid})
```

- **@RequestParam**

```java
通过@RequestParam，例如blogs?blogId=1
```

