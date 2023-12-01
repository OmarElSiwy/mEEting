# mEEting - Electrical Engineering Forum App

Welcome to **mEEting**, your go-to application for connecting with the Electrical Engineering community and staying updated on the latest news and academic discussions in the field. This mobile app, compatible with both iOS and Android devices, is designed to make it easier for EE enthusiasts to engage in meaningful conversations, access academic resources, and stay informed about the latest developments in Electrical Engineering.

## Table of Contents
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Getting Started](#getting-started)
- [Current Progress](#current-progress)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- **Forum for EE Enthusiasts**: Join a vibrant community of Electrical Engineering enthusiasts and professionals to discuss topics related to EE, academic resources, and the latest news.

- **Connect with Academic Representatives**: Stay updated with academic representatives and their latest announcements directly on the app.

- **Integration with Uwaterloo API**: Seamlessly integrate with the Uwaterloo API to access valuable educational resources and information.

- **Native iOS and Android Compatibility**: Utilize the app on both iOS and Android platforms for a consistent user experience.

- **Real-Time Discussions**: Engage in real-time discussions and share your insights with fellow engineers.

- **Offline Access**: Access content and discussions offline, ensuring that you stay connected even when you don't have an internet connection.

---

## Technologies Used

This project is built using a variety of technologies to provide a robust and efficient experience:

- **React Native**: The framework for building the cross-platform mobile app.

- **Three.js**: Used for 3D graphics rendering.

- **Go (Backend)**: Powers the backend of the application.

- **Azure SQL**: The database system for storing and managing data.

- **Uwaterloo API**: Integrates with the Uwaterloo API for educational resources and information.

- **Docker**: Includes a Dockerfile for testing and replicating the development environment.

- **Native Modules (Java & Swift)**: Interfaces natively with Android and iOS modules, connecting them to TypeScript for seamless API integration without relying solely on internet connections.

---

## Installation

To set up the development environment for mEEting, follow these steps:

1. Clone this repository to your local machine.

2. Build the Docker container using the provided Dockerfile.

3. Run the Expo development server to start the application.

4. Use the provided `Expotest.sh` script to streamline the build and execution process.

---

## Getting Started

1. After setting up the development environment, run the Expo development server:
```
npx expo start
```

2. The app will be available on ports 19000 (Expo) and 8081 (React Native).

3. Connect your Android or iOS device to test the app, or use an emulator/simulator for testing.

4. Start exploring the world of Electrical Engineering with mEEting!

---

## Current Progress
1. Setup Dockerfile Environment:

Created a Dockerfile to establish a standardized development environment for cross-collaboration.
Developed an executable (Expotest.sh) to streamline code execution within the virtual desktop.
✅ Completed.

2. Backend Environment Setup:

Configured the backend environment to handle Posting, Getting, and Querying of information.
Utilized a custom database setup for enhanced security compared to external APIs.
✅ Completed.

3. User Authentication:

Implemented a robust login and registration system with verification specific to University of Waterloo Electrical Engineering students.
✅ Completed.

4. Frontend Dependencies:

Installed all required dependencies for the frontend project to ensure smooth development.
✅ Completed.

5. Native Module Integration:

Integrated the external backend code with the frontend using native modules.
Developed Swift and Java interface code to facilitate communication with the GO code, enabling seamless interaction with TypeScript.
✅ Completed.

6. Frontend Design:

Work in progress: Currently designing an aesthetically pleasing frontend to engage and attract students, with the aim of surpassing the appeal of mainstream social media platforms like Instagram.

ISSUES: No MacOS to test IOS version on (Buying a MacBook soon)

## Contributing

We welcome contributions from the community to make mEEting even better. If you'd like to contribute, please follow our [Contribution Guidelines](CONTRIBUTING.md).

---

## License

This project is licensed under the [MIT License](LICENSE.md).

---

Thank you for choosing mEEting for all your Electrical Engineering discussions and news updates. We look forward to building a vibrant and supportive community together. If you have any questions or suggestions, please don't hesitate to reach out!


# Troubleshooting

If you can't get this to work, see the [Troubleshooting](https://reactnative.dev/docs/troubleshooting) page.

# Learn More

To learn more about React Native, take a look at the following resources:

- [React Native Website](https://reactnative.dev) - learn more about React Native.
- [Getting Started](https://reactnative.dev/docs/environment-setup) - an **overview** of React Native and how setup your environment.
- [Learn the Basics](https://reactnative.dev/docs/getting-started) - a **guided tour** of the React Native **basics**.
- [Blog](https://reactnative.dev/blog) - read the latest official React Native **Blog** posts.
- [`@facebook/react-native`](https://github.com/facebook/react-native) - the Open Source; GitHub **repository** for React Native.
