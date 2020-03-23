# Use AdoptOpenJDK for base image.
FROM registry.cn-hangzhou.aliyuncs.com/knative-sample/openjdk:8-alpine

# Copy the jar to the production image from the builder stage.
COPY  seckill-order-service-1.0.0.jar /app/seckill-order-service-1.0.0.jar

# Run the web service on container startup.
CMD ["java", "-jar", "-DenvType=aliyun", "/app/seckill-order-service-1.0.0.jar"]
