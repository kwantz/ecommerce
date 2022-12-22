FROM openjdk:17-alpine

WORKDIR /app

# Only copy dependency-related files
COPY product/gradle /app/gradle/
COPY product/build.gradle /app/
COPY product/gradlew /app/
COPY product/gradlew.bat /app/
COPY product/settings.gradle /app/

# Only download dependencies
# Eat the expected build failure since no source code has been copied yet
RUN ./gradlew clean build --no-daemon > /dev/null 2>&1 || true

# Copy all files
COPY product/ /app/

# Do the actual build
RUN ./gradlew clean build --no-daemon

ENTRYPOINT ["java", "-jar", "/app/build/libs/product-0.0.1-SNAPSHOT.jar"]