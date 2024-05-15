[![Go](https://github.com/Kontentski/weatherly/actions/workflows/go.yml/badge.svg)](https://github.com/Kontentski/weatherly/actions/workflows/go.yml)

# Weatherly API

Weatherly is a weather API written in Go, designed to provide weather updates directly via the command line or through a browser. The API responds with weather information for the current day based on the IP address of the request. It supports output in both HTML and plain text formats, automatically adjusting based on the request headers.

## Features

- **Auto-detect Response Format**: Automatically detects whether the client supports HTML. If not, it responds with plain text suitable for terminal use.
- **IP-based Localization**: Determines the weather based on the IP address from which the request is made, providing localized weather information.
- **Docker Support**: Packaged as a Docker container for easy deployment and scalability.
- **Cloud Deployment**: Currently hosted on Google Cloud, accessible anytime from anywhere.

## How to Use

You can access the Weatherly API through the following link: [Weatherly API](https://weatherly-4zvrwzl5sq-uc.a.run.app/)

<img width="617" alt="image" src="https://github.com/Kontentski/weatherly/assets/150854976/2551a213-6b38-4ba5-9285-a5371ac846ae">

### Using Curl

To fetch weather information from your terminal, you can use `curl`:

```bash
curl https://weatherly-4zvrwzl5sq-uc.a.run.app/
```
For ease of use, you can add the following function to your shell profile (e.g., .bashrc or .zshrc):
```
weatherly() {
    curl "https://weatherly-4zvrwzl5sq-uc.a.run.app/$1"
}
```
After adding this function, you can simply call weatherly from your terminal to get the weather information.

<img width="419" alt="image" src="https://github.com/Kontentski/weatherly/assets/150854976/11b694e1-a8ff-487c-a134-8e66c2fe2f29">

# Installation

To run Weatherly locally, ensure you have Docker installed. Then, follow these steps:

- **Clone the repository**:
  ```bash
  git clone https://github.com/Kontentski/weatherly
  ```
- **Obtain an API key from [WeatherAPI.com](https://www.weatherapi.com/)**:
  
  Sign up for a free account and get your API key.
  
- **Build the Docker image**:
  ```bash
  docker build --tag weatherly .
  ```
- **Run the container**:
  ```bash
  docker run -p 8080:8080 -e KEY=YourAPIKeyHere weatherly
  ```
After running these steps, Weatherly will be accessible on localhost:8080.

# Contributions

Contributions are welcome! If you'd like to contribute to the project, please fork the repository and submit a pull request.

# License

Weatherly is released under the MIT License. See the LICENSE file for more details.




