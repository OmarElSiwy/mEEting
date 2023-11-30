# Use an official OpenJDK image with Java 11
FROM openjdk:8-jdk

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy the current directory contents into the container
COPY . .

# Install Node.js and other dependencies
RUN apt-get update -qq && apt-get install -y nodejs npm unzip

# Install necessary libraries for XML binding
RUN apt-get update -qq && apt-get install -y libjaxb-api-java

# Install any needed packages specified in package.json
RUN npm install

# Install Go
RUN wget https://golang.org/dl/go1.18.linux-amd64.tar.gz
RUN tar -xvf go1.18.linux-amd64.tar.gz
RUN mv go /usr/local
ENV PATH="/usr/local/go/bin:${PATH}"

# Set GOPATH
ENV GOPATH="/root/go"
ENV PATH="${GOPATH}/bin:${PATH}"

# Download and setup Android SDK
ENV ANDROID_HOME="/usr/local/android-sdk"
ENV PATH="$PATH:${ANDROID_HOME}/tools:${ANDROID_HOME}/tools/bin:${ANDROID_HOME}/platform-tools"
RUN mkdir -p "$ANDROID_HOME" && cd "$ANDROID_HOME" && \
    wget -q https://dl.google.com/android/repository/sdk-tools-linux-4333796.zip -O android-sdk-tools.zip && \
    unzip -q android-sdk-tools.zip && \
    rm android-sdk-tools.zip

# Accept licenses
RUN yes | sdkmanager --licenses

# Install necessary Android SDK packages
RUN sdkmanager "platform-tools" "platforms;android-29" "build-tools;29.0.2"

# Install the Android NDK
RUN sdkmanager "ndk;21.3.6528147"  # Replace with the version you need

# Install GoMobile
RUN go install golang.org/x/mobile/cmd/gomobile@latest
RUN gomobile init

# Copy the backend directory and build GoMobile bindings
COPY backend /usr/src/app/backend
WORKDIR /usr/src/app/backend
RUN go install golang.org/x/mobile/cmd/gomobile@latest
RUN gomobile bind -target=android -o /usr/src/app/android/app/libs/backend.aar .
RUN gomobile bind -target=ios -o /usr/src/app/ios/Backend.framework .

# Switch back to the main app directory
WORKDIR /usr/src/app

# Make ports 19000 (Expo) and 8081 (React Native) available to the world outside this container
EXPOSE 19000
EXPOSE 8081

# Run Expo start 
# CMD ["expo", "start"]