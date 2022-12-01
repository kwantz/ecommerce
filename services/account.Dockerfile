FROM openjdk:17-alpine

WORKDIR /app

# Only copy dependency-related files
COPY account/gradle /app/gradle/
COPY account/build.gradle /app/
COPY account/gradlew /app/
COPY account/gradlew.bat /app/
COPY account/settings.gradle /app/

# Only download dependencies
# Eat the expected build failure since no source code has been copied yet
RUN ./gradlew clean build --no-daemon > /dev/null 2>&1 || true

# Copy all files
COPY account/ /app/

# Do the actual build
RUN ./gradlew clean build --no-daemon

ENTRYPOINT ["java", "-jar", "/app/build/libs/account-0.0.1-SNAPSHOT.jar"]