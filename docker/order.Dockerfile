FROM openjdk:17-alpine

WORKDIR /app

# Only copy dependency-related files
COPY order/gradle /app/gradle/
COPY order/build.gradle /app/
COPY order/gradlew /app/
COPY order/gradlew.bat /app/
COPY order/settings.gradle /app/

# Only download dependencies
# Eat the expected build failure since no source code has been copied yet
RUN ./gradlew clean build --no-daemon > /dev/null 2>&1 || true

# Copy all files
COPY order/ /app/

# Do the actual build
RUN ./gradlew clean build --no-daemon

ENTRYPOINT ["java", "-jar", "/app/build/libs/order-0.0.1-SNAPSHOT.jar"]